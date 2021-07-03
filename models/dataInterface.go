package models

type Store interface {
	AddScore(postID, userID string, score int)
	AddPost(name, path, userID string) *Post
	Follow(followerID, followingID string)
	Unfollow(followerID, followingID string)

	GetScores() []*Score
	GetPosts() []*Post
	GetPostFromID(id string) (*Post, error)
	GetFollowRecords() []FollowRecord
	GetUserByID(id string) (*User, error)
}

var OnlyFrogsSession Store

func NewOnlyFrogsSession(s Store) {
	OnlyFrogsSession = s
}
