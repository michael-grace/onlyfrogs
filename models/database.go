package models

type Database struct {
}

func (d *Database) AddScore(postID, userID string, score int) {

}

func (d *Database) AddPost(name, path, userID string) *Post {
	return nil
}

func (d *Database) Follow(followerID, followingID string) {

}

func (d *Database) Unfollow(followerID, followingID string) {

}

func (d *Database) GetPosts() []*Post {
	return nil
}

func (d *Database) GetPostFromID(id string) (*Post, error) {
	return nil, nil
}

func (d *Database) GetTotalScoreForPost(id string) int {
	return 0
}

func (d *Database) GetUserByID(id string) (*User, error) {
	return nil, nil
}

func (d *Database) GetUserFollowing(id string) []*User {
	return nil
}

func (d *Database) GetUserFollowers(id string) []*User {
	return nil
}
