package models

type FollowRecord struct {
	FollowerID  string `json:"followerID"`
	FollowingID string `json:"followingID"`
}
