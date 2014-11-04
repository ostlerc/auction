package auction

const (
	Slow = iota
	Fast
	Random
	True

	slowstr   = "slow"
	faststr   = "fast"
	randomstr = "random"
	truestr   = "true"
)

type Bidder interface {
	Bid(int) *Bid
}

type RangeBidder struct {
	m int
	t string
	i int
}

type CustomBidder struct {
	money int
	max   int
	i     int
}

func SlowBidder(n int) Bidder {
	return &RangeBidder{
		m: n / 5,
		t: slowstr,
		i: n,
	}
}

func FastBidder(n int) Bidder {
	return &RangeBidder{
		m: n / 3,
		t: faststr,
		i: n,
	}
}

func RandomBidder(n int) Bidder {
	return &RangeBidder{
		m: R.Intn(n / 2),
		t: randomstr,
		i: n,
	}
}

func TrueBidder(money, i, price int) Bidder {
	return &CustomBidder{
		money: money,
		i:     i,
		max:   R.Intn(price) + 1,
	}
}

func (r *RangeBidder) Bid(n int) *Bid {
	res := make([]int, n)
	high := r.i
	for i := 0; i < n; i++ {
		if r.m > 0 {
			high -= R.Intn(r.m)
		}
		if high < 0 {
			high = 0
		}
		res[i] = high
	}
	return &Bid{
		Type: r.t,
		Bids: res,
	}
}

func (t *CustomBidder) Bid(n int) *Bid {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = t.money / (t.max * (i + 1))
		if res[i] > t.i {
			res[i] = t.i
		}
	}
	return &Bid{
		Type: truestr,
		Bids: res,
	}
}
