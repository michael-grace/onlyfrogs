package models

import "time"

type Post struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	PhotoPath string    `json:"photo"`
	UserID    string    `json:"userID"`
	Time      time.Time `json:"time"`
}
