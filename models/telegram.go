package models

type Message struct {
	ChatID    int64               `json:"chat_id"`
	Token     string              `json:"token"`
	Image     string              `json:"image"`
	Text      string              `json:"text"`
	EffectID  string              `json:"effect_id"`
	Keyboard  []map[string]string `json:"keyboard"`
}
