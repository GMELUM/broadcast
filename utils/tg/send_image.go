package tg

import (
	"broadcast/models"

	"github.com/NicoNex/echotron/v3"
)

func SendImage(
	api *echotron.API,
	task *models.Message,
	chat_id int64,
) (res echotron.APIResponseMessage, err error) {

	var img = echotron.NewInputFileURL(task.Image)

	var keyboard [][]echotron.InlineKeyboardButton
	for _, row := range task.Keyboard {
		var buttons []echotron.InlineKeyboardButton
		for text, url := range row {
			buttons = append(buttons, echotron.InlineKeyboardButton{
				Text: text,
				URL:  url,
			})
		}
		keyboard = append(keyboard, buttons)
	}

	return api.SendPhoto(
		img,
		chat_id,
		&echotron.PhotoOptions{
			MessageEffectID: task.EffectID,
			ReplyMarkup: echotron.InlineKeyboardMarkup{
				InlineKeyboard: keyboard,
			},
			Caption: task.Text,
		},
	)

}
