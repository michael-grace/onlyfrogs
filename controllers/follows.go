package controllers

import (
	"encoding/json"
	"net/http"

	m "github.com/michael-grace/onlyfrogs/models"
)

func Follow(w http.ResponseWriter, r *http.Request) {
	var newFollow m.FollowRecord
	err := json.NewDecoder(r.Body).Decode(&newFollow)
	if err != nil {
		panic(err) // TODO
	}
	m.OnlyFrogsSession.Follow(newFollow.FollowerID, newFollow.FollowingID)
	jsonData, err := json.Marshal(newFollow)
	if err != nil {
		panic(err) // TODO
	}
	w.Write(jsonData)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	var newUnfollow m.FollowRecord
	err := json.NewDecoder(r.Body).Decode(&newUnfollow)
	if err != nil {
		panic(err) // TODO
	}
	m.OnlyFrogsSession.Unfollow(newUnfollow.FollowerID, newUnfollow.FollowingID)
	jsonData, err := json.Marshal(newUnfollow)
	if err != nil {
		panic(err) // TODO
	}
	w.Write(jsonData)
}
