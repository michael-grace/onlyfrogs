package data

import (
	m "github.com/michael-grace/onlyfrogs/models"
)

type TestSuite struct {
	Users  []*m.User
	Posts  []*m.Post
	Scores []*m.Score
}

func (t *TestSuite) GetScoreForPost(postID int) int {
	postScore := 0
	for _, score := range t.Scores {
		if score.PostID == postID {
			postScore += score.Score
		}
	}
	return postScore
}

func (t *TestSuite) AddScoreToPost(postID, userID, score int) {
	t.Scores = append(t.Scores, &m.Score{UserID: userID, PostID: postID, Score: score})
}
