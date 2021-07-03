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

	if OnlyFrogsSession.GetTotalScoreForPost(post.ID) != 2 {
		t.Errorf("Expected Score 2, got score %v", OnlyFrogsSession.GetTotalScoreForPost(post.ID))
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

	origScore := OnlyFrogsSession.GetTotalScoreForPost(post.ID)
	OnlyFrogsSession.AddScore("1", "3", 5)
	newScore := OnlyFrogsSession.GetTotalScoreForPost(post.ID)

	if newScore != origScore+5 {
		t.Errorf("Expected Score %v, got score %v", origScore+5, newScore)
	}
}
