package utils

import (
	"math/rand"
	"time"
)

func NewRand() *rand.Rand {
	seed := time.Now().UnixNano()
	s := rand.NewSource(seed)
	return rand.New(s)
}

func Floatn(r *rand.Rand, minFloat float64, maxFloat float64) float64 {
	return minFloat + r.Float64()*(maxFloat-minFloat)
}
