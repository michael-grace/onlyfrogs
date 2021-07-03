package models

type Store interface {
	AddScore(postID, userID string, score int)
	AddPost(name, path, userID string) *Post

	GetScores() []*Score
	GetPosts() []*Post
	GetPostFromID(id string) (*Post, error)
}

var OnlyFrogsSession Store

func NewOnlyFrogsSession(s Store) {
	OnlyFrogsSession = s
}
