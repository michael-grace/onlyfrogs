package models

import (
	"fmt"

	"github.com/google/uuid"
)

type TestSuite struct {
	Users  []*User
	Posts  []*Post
	Scores []*Score
}

func (t *TestSuite) GetScores() []*Score {
	return t.Scores
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

func SetUp() {
	NewOnlyFrogsSession(&TestSuite{
		Users: []*User{
			&User{
				ID:       "1",
				Username: "username1",
			},
			&User{
				ID:       "2",
				Username: "username2",
			},
		},
		Posts: []*Post{
			&Post{
				ID:        "1",
				Name:      "Frog One",
				PhotoPath: "1.jpg",
				UserID:    "1",
			},
			&Post{
				ID:        "2",
				Name:      "Frog Two",
				PhotoPath: "2.jpg",
				UserID:    "2",
			},
		},
		Scores: []*Score{
			&Score{
				UserID: "1",
				PostID: "2",
				Score:  4,
			},
			&Score{
				UserID: "2",
				PostID: "1",
				Score:  -1,
			},
			&Score{
				UserID: "1",
				PostID: "1",
				Score:  3,
			},
		},
	})
}
