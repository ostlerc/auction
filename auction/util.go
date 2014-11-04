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

//Average all except at index n. Regular average if n == -1
func (i Ints) Ave(n int) float64 {
	sum := i.Sum()
	if n != -1 {
		sum -= i[n]
		return float64(sum) / float64(len(i)-1)
	}
	return float64(sum) / float64(len(i))
}

func (i Ints) Add(r Ints) {
	if len(i) != len(r) {
		panic("incorrect adding")
	}
	for x := 0; x < len(i); x++ {
		i[x] += r[x]
	}
}

func (i Ints) Div(n int) []float64 {
	res := make([]float64, len(i))
	for x := 0; x < len(i); x++ {
		res[x] = float64(i[x]) / float64(n)
	}
	return res
}
