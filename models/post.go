package models

type Post struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	PhotoPath string `json:"photo"`
	UserID    string `json:"userID"`
}

// This shouldn't be here, and should be on the Store interface

func (p Post) GetScore() int {
	totalScore := 0
	for _, score := range OnlyFrogsSession.GetScores() {
		if score.PostID == p.ID {
			totalScore += score.Score
		}
	}
	return totalScore
}
