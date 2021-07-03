package controllers

import (
	"encoding/json"
	"net/http"

	m "github.com/michael-grace/onlyfrogs/models"
)

// TODO User Tokens
func AddScore(w http.ResponseWriter, r *http.Request) {
	var newScore m.Score
	err := json.NewDecoder(r.Body).Decode(&newScore)
	if err != nil {
		panic(err) // TODO
	}

	m.OnlyFrogsSession.AddScore(newScore.PostID, newScore.UserID, newScore.Score)
	jsonData, err := json.Marshal(newScore)
	if err != nil {
		panic(err) // TODO
	}
	w.Write(jsonData)
}
