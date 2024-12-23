package utils

import (
	"time"

	"github.com/schollz/progressbar/v3"
)

func NewProgress(title string, count int) *progressbar.ProgressBar {
	return progressbar.NewOptions(count,
		progressbar.OptionSetDescription(title),
		progressbar.OptionShowCount(),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionThrottle(time.Millisecond * 100),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "█",
			SaucerHead:    "█",
			SaucerPadding: " ",
			BarStart:      "|",
			BarEnd:        "|",
		}),
	)
}
