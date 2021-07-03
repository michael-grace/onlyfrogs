package models

type Store interface {
	AddScore(postID, userID string, score int)
	AddPost(name, path, userID string) *Post
	Follow(followerID, followingID string) error
	Unfollow(followerID, followingID string) error

	GetPosts() []*Post
	GetPostFromID(id string) (*Post, error)
	GetTotalScoreForPost(id string) int

	GetUserByID(id string) (*User, error)
	GetUserFollowing(id string) []*User
	GetUserFollowers(id string) []*User
	GetFeedForUser(id string) []*Post
}

var OnlyFrogsSession Store

func NewOnlyFrogsSession(s Store) {
	OnlyFrogsSession = s
}
