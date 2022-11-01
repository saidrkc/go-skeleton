package inmemory

import "math/rand"

func UserScoreRandom(i int) UserScoreInMemory {
	return UserScoreInMemory{
		UserId: i,
		Total:  rand.Float32(),
	}
}
