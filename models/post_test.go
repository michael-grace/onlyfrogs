package models

import (
	"testing"
)

func TestAddPost(t *testing.T) {
	if OnlyFrogsSession == nil {
		SetUp()
	}
	origLen := len(OnlyFrogsSession.GetPosts())
	OnlyFrogsSession.AddPost("New Frog", "new.png", "2")
	newLen := len(OnlyFrogsSession.GetPosts())
	if newLen != (origLen + 1) {
		t.Errorf("Expected Length %v, got length %v", origLen+1, newLen)
	}
}

func TestGetPostScore(t *testing.T) {
	if OnlyFrogsSession == nil {
		SetUp()
	}

	post, err := OnlyFrogsSession.GetPostFromID("1")
	if err != nil {
		t.Errorf(err.Error())
	}

	if post.GetScore() != 2 {
		t.Errorf("Expected Score 2, got score %v", post.GetScore())
	}

}

func TestChangePostScore(t *testing.T) {
	if OnlyFrogsSession == nil {
		SetUp()
	}

	post, err := OnlyFrogsSession.GetPostFromID("1")
	if err != nil {
		t.Errorf(err.Error())
	}

	origScore := post.GetScore()
	OnlyFrogsSession.AddScore("1", "3", 5)
	newScore := post.GetScore()

	if newScore != origScore+5 {
		t.Errorf("Expected Score %v, got score %v", origScore+5, newScore)
	}
}
