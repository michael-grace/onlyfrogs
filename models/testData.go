package models

import (
	"fmt"

	"github.com/google/uuid"
)

type TestSuite struct {
	Users         []*User
	Posts         []*Post
	Scores        []*Score
	FollowRecords []FollowRecord
}

func (t *TestSuite) GetPosts() []*Post {
	return t.Posts
}

func (t *TestSuite) GetPostFromID(id string) (*Post, error) {
	for _, post := range t.Posts {
		if post.ID == id {
			return post, nil
		}
	}
	return nil, fmt.Errorf("Post with ID %v not found", id)
}

func (t *TestSuite) GetTotalScoreForPost(id string) int {
	var totalScore int
	for _, score := range t.Scores {
		if score.PostID == id {
			totalScore += score.Score
		}
	}
	return totalScore
}

func (t *TestSuite) GetUserByID(id string) (*User, error) {
	for _, user := range t.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("User with ID %v not found", id)
}

func (t *TestSuite) GetUserFollowing(id string) []*User {
	var following []*User
	for _, record := range t.FollowRecords {
		if record.FollowerID == id {
			user, err := t.GetUserByID(record.FollowingID)
			if err != nil {
				panic(err) // TODO
			}
			following = append(following, user)
		}
	}
	return following
}

func (t *TestSuite) GetUserFollowers(id string) []*User {
	var followers []*User
	for _, record := range t.FollowRecords {
		if record.FollowingID == id {
			user, err := t.GetUserByID(record.FollowerID)
			if err != nil {
				panic(err) // TODO
			}
			followers = append(followers, user)
		}
	}
	return followers
}

func (t *TestSuite) AddScore(postID, userID string, score int) {
	t.Scores = append(t.Scores, &Score{UserID: userID, PostID: postID, Score: score})
}

func (t *TestSuite) AddPost(name, path, userID string) *Post {
	newPost := &Post{
		ID:        uuid.NewString(),
		Name:      name,
		PhotoPath: path,
		UserID:    userID,
	}
	t.Posts = append(t.Posts, newPost)
	return newPost
}

func (t *TestSuite) Follow(followerID, followingID string) error {
	if followerID == followingID {
		return fmt.Errorf("you can't follow yourself")
	}
	t.FollowRecords = append(t.FollowRecords, FollowRecord{
		FollowerID:  followerID,
		FollowingID: followingID,
	})
	return nil
}

func (t *TestSuite) Unfollow(followerID, followingID string) error {
	if followerID == followingID {
		return fmt.Errorf("you can't unfollow yourself")
	}
	var recordIdx int = -1
	for idx, val := range t.FollowRecords {
		if val.FollowerID == followerID && val.FollowingID == followingID {
			recordIdx = idx
			break
		}
	}
	t.FollowRecords = append(t.FollowRecords[:recordIdx], t.FollowRecords[recordIdx+1:]...)
	return nil
}

func SetUp() {
	NewOnlyFrogsSession(&TestSuite{
		Users: []*User{
			{
				ID:       "1",
				Username: "username1",
			},
			{
				ID:       "2",
				Username: "username2",
			},
		},
		Posts: []*Post{
			{
				ID:        "1",
				Name:      "Frog One",
				PhotoPath: "1.jpg",
				UserID:    "1",
			},
			{
				ID:        "2",
				Name:      "Frog Two",
				PhotoPath: "2.jpg",
				UserID:    "2",
			},
		},
		Scores: []*Score{
			{
				UserID: "1",
				PostID: "2",
				Score:  4,
			},
			{
				UserID: "2",
				PostID: "1",
				Score:  -1,
			},
			{
				UserID: "1",
				PostID: "1",
				Score:  3,
			},
		},
		FollowRecords: []FollowRecord{
			{
				FollowerID:  "1",
				FollowingID: "2",
			},
		},
	})
}
