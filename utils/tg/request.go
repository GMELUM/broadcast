package tg

import (
	"broadcast/models"

	"github.com/NicoNex/echotron/v3"
)

var cache = make(map[string]*echotron.API)

func Request(
	task *models.Message,
	token string,
	chat_id int64,
) (res echotron.APIResponseMessage, err error) {

	a, ok := cache[token]
	if !ok {
		api := echotron.NewAPI(token)
		a = &api
		cache[token] = &api
	}

	if task.Image != "" {
		return SendImage(a, task, chat_id)
	}

	return SendMessage(a, task, chat_id)

}
