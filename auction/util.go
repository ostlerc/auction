package auction

import (
	"math/rand"
	"time"
)

var (
	R = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func ArEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
