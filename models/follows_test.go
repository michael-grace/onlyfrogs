package models

import (
	"testing"
)

func TestFollow(t *testing.T) {
	if OnlyFrogsSession == nil {
		SetUp()
	}
	origLenUser1 := len(OnlyFrogsSession.GetUserFollowers("1"))
	origLenUser2 := len(OnlyFrogsSession.GetUserFollowing("2"))
	OnlyFrogsSession.Follow("2", "1")
	newLenUser1 := len(OnlyFrogsSession.GetUserFollowers("1"))
	newLenUser2 := len(OnlyFrogsSession.GetUserFollowing("2"))
	if newLenUser1 != origLenUser1+1 || newLenUser2 != origLenUser2+1 {
		t.Errorf("Failed Follow Test. User 1 was %v, should now be %v but was %v. User 2 was %v, and should now be %v but was %v",
			origLenUser1,
			origLenUser1+1,
			newLenUser1,
			newLenUser2,
			origLenUser2+1,
			newLenUser2)
	}
}

func TestUnfollow(t *testing.T) {
	if OnlyFrogsSession == nil {
		SetUp()
	}
	origLenUser1 := len(OnlyFrogsSession.GetUserFollowing("1"))
	origLenUser2 := len(OnlyFrogsSession.GetUserFollowers("2"))
	OnlyFrogsSession.Unfollow("1", "2")
	newLenUser1 := len(OnlyFrogsSession.GetUserFollowing("1"))
	newLenUser2 := len(OnlyFrogsSession.GetUserFollowers("2"))
	if newLenUser1 != origLenUser1-1 || newLenUser2 != origLenUser2-1 {
		t.Errorf("Failed Unfollow Test. User 1 was %v, should now be %v but was %v. User 2 was %v, and should now be %v but was %v",
			origLenUser1,
			origLenUser1-1,
			newLenUser1,
			newLenUser2,
			origLenUser2-1,
			newLenUser2)
	}
}

func TestFollowYourself(t *testing.T) {
	if OnlyFrogsSession == nil {
		SetUp()
	}
	err := OnlyFrogsSession.Follow("1", "1")
	if err == nil {
		t.Error("failed to catch following yourself")
	}
}

func TestUnfollowYourself(t *testing.T) {
	if OnlyFrogsSession == nil {
		SetUp()
	}
	err := OnlyFrogsSession.Unfollow("1", "1")
	if err == nil {
		t.Error("failed to catch unfollowing yourself")
	}
}
