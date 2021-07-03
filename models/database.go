package models

type Database struct {
}

func (d *Database) AddScore(postID, userID string, score int) {

}

func (d *Database) AddPost(name, path, userID string) *Post {
	return nil
}

func (d *Database) GetScores() []*Score {
	return nil
}

func (d *Database) GetPosts() []*Post {
	return nil
}

func (d *Database) GetPostFromID(id string) (*Post, error) {
	return nil, nil
}
