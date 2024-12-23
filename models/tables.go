package models

type Queue struct {
	ID     int  `json:"id"`
	User   int  `json:"user"`
	Status int `json:"success"`
}
