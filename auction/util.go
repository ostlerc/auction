package auction

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	R = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Ints []int

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

func (i Ints) String() []string {
	res := make([]string, len(i))
	for i, v := range i {
		res[i] = strconv.Itoa(v)
	}
	return res
}

func (i Ints) Sum() int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
