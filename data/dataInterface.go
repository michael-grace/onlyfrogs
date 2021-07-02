package data

type Store interface {
	AddScoreToPost(int, int, int)
	GetScoreForPost(int) int
}
