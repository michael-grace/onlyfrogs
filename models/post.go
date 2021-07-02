package models

type Post struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PhotoPath string `json:"photo"`
	UserID    int    `json:"userID"`
}
