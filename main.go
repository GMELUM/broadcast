package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"broadcast/config"
	"broadcast/models"
	"broadcast/storage"
	"broadcast/utils"
	"broadcast/utils/cfg"
	ctx "broadcast/utils/context"
	"sync"

	"broadcast/utils/sqlite"
	"broadcast/utils/tg"
	"os"
	"time"

	"github.com/paulbellamy/ratecounter"
	"golang.org/x/time/rate"
)

func main() {
	// Handle graceful shutdown and panic recovery
	defer utils.HandlerExit()

	// Process configuration and write settings
	err := cfg.ConfigWriting()
	if err != nil {
		println(err.Error())
		return
	}

	// Initialize SQLite database connection
	db, err := sqlite.New("./output.db")
	if err != nil {
		println(err.Error())
		return
	}

	// Set database connection in global context
	ctx.SetSQLite(db)

	// Check if queue table exists
	exists, err := storage.TableExists("queue")
	if err != nil {
		println(err.Error())
		return
	}

	// Create and populate queue table if it doesn't exist
	if !exists {
		// Find CSV file in current directory
		path, err := utils.SearchCSV()
		if err != nil {
			println(err.Error())
			return
		}

		// Open CSV file
		file, err := os.Open(path)
		if err != nil {
			println("failed to open CSV file")
			return
		}

		// Create CSV reader and read all records
		reader := csv.NewReader(bufio.NewReader(file))
		records, err := reader.ReadAll()
		if err != nil {
			println("failed to read CSV file")
			return
		}

		// Validate CSV has at least one data row
		if len(records) < 2 {
			println("CSV file is empty")
			return
		}

		// Create queue table in database
		err = storage.CreateTable()
		if err != nil {
			println(err.Error())
			return
		}

		// Insert CSV data into queue table in batches
		err = storage.FillOutTable("queue", records, config.SplitChunkCount)
		if err != nil {
			println(err.Error())
			return
		}
	}

	// Get count of unprocessed rows
	allCount, err := storage.CountRowsWithStatus(0)
	if err != nil {
		println(err.Error())
		return
	}

	// Initialize rate counter and progress bar
	per := ratecounter.NewRateCounter(time.Second)
	bar := utils.NewProgress(fmt.Sprintf("[%-9v] %5vrps  ", "broadcast", per.Rate()), allCount)

	// Read message template from JSON file
	message, err := utils.ReadDataMessage("./msg.json")
	if err != nil {
		println(err.Error())
		return
	}

	// Create rate limiter for API requests
	limiter := rate.NewLimiter(rate.Limit(config.RPSLimit), 1)

	// Initialize wait group and context
	var wg sync.WaitGroup
	var c = context.Background()

	page := 0

	// Process users in batches
	for {
		// Get batch of users
		users, err := storage.GetUsers(page, 100, 1)
		if err != nil {
			println(err.Error())
			return
		}

		// Exit if no more users
		if len(*users) == 0 {
			break
		}

		// Process each user in batch
		for _, user := range *users {
			// Skip already processed users
			if user.Status != 0 {
				per.Incr(1)
				bar.Add(1)
				bar.Describe(fmt.Sprintf("[%-9v] %5vrps  ", "broadcast", per.Rate()))
				continue
			}

			// Apply rate limiting
			err := limiter.Wait(c)
			if err != nil {
				println("Error in rate limiter:", err)
				continue
			}

			// Start goroutine for API request
			wg.Add(1)
			go func(u models.Queue) {
				defer wg.Done()

				// Make API request to send message
				res, err := tg.Request(
					message,
					config.Token,
					int64(user.User),
				)

				// Update progress tracking
				per.Incr(1)
				bar.Add(1)
				bar.Describe(fmt.Sprintf("[%-9v] %5vrps  ", "broadcast", per.Rate()))

				// Handle request error
				if err != nil {
					err = storage.UpdateStatusByID(user.ID, 2)
					if err != nil {
						println(err.Error())
						return
					}
					return
				}

				// Update status on successful request
				if res.Ok {
					err = storage.UpdateStatusByID(user.ID, 1)
					if err != nil {
						println(err.Error())
						return
					}
				}

				// Log API errors
				if !res.Ok {
					println(res.APIResponseBase.ErrorCode)
				}

			}(user)
		}

		page++
	}

	// Wait for all goroutines to complete
	wg.Wait()

}
