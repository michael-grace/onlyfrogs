package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/michael-grace/onlyfrogs/models"
)

type PostWithScore struct {
	m.Post
	Score int `json:"score"`
}

// TODO User Tokens
func AddPost(w http.ResponseWriter, r *http.Request) {
	var newPost m.Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		panic(err) // TODO
	}
	createdPost := m.OnlyFrogsSession.AddPost(newPost.Name, newPost.PhotoPath, newPost.UserID)
	jsonData, err := json.Marshal(createdPost)
	if err != nil {
		panic(err) // TODO
	}
	w.Write(jsonData)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["uuid"]
	if !ok {
		panic("AAHH") // TODO
	}
	post, err := m.OnlyFrogsSession.GetPostFromID(id)
	if err != nil {
		panic(err) // TODO
	}
	jsonData, err := json.Marshal(&PostWithScore{Post: *post, Score: m.OnlyFrogsSession.GetTotalScoreForPost(post.ID)})
	if err != nil {
		panic(err) // TODO
	}
	w.Write(jsonData)

}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	allPosts := m.OnlyFrogsSession.GetPosts()
	var allPostsScores []*PostWithScore
	for _, val := range allPosts {
		allPostsScores = append(allPostsScores, &PostWithScore{Post: *val, Score: m.OnlyFrogsSession.GetTotalScoreForPost(val.ID)})
	}
	jsonData, err := json.Marshal(allPostsScores)
	if err != nil {
		panic(err) // TODO
	}
	w.Write(jsonData)
}
