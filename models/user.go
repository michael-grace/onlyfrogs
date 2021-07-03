package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// Tbh, these functions shouldn't be here, and should be straight on
// the Store interface

func (u *User) GetFollowing() []*User {
	var following []*User
	for _, record := range OnlyFrogsSession.GetFollowRecords() {
		if record.FollowerID == u.ID {
			user, err := OnlyFrogsSession.GetUserByID(record.FollowingID)
			if err != nil {
				panic(err) // TODO
			}
			following = append(following, user)
		}
	}
	return following
}

func (u *User) GetFollowers() []*User {
	var followers []*User
	for _, record := range OnlyFrogsSession.GetFollowRecords() {
		if record.FollowingID == u.ID {
			user, err := OnlyFrogsSession.GetUserByID(record.FollowerID)
			if err != nil {
				panic(err)
			}
			followers = append(followers, user)
		}

	}
	return followers
}
