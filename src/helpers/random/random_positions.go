package random

import (
	"math/rand"
	"time"
)

func GetRandomPositions(size int, n uint) []int {
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(size)

	var positions []int
	positions = append(positions, p[:n]...)

	return positions
}
