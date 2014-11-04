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

func TrueBidder(money, max int) Bidder {
	return &CustomBidder{
		money: money,
		max:   max,
	}
}

func (r *RangeBidder) Bid(n int) *Bid {
	res := make([]int, n)
	high := r.i
	for i := 0; i < n; i++ {
		high -= R.Intn(r.m)
		if high < 0 {
			high = 0
		}
		res[i] = high
	}
	return &Bid{
		BidType: r.t,
		Bids:    res,
	}
}

func (t *CustomBidder) Bid(n int) *Bid {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = t.money / (i + 1)
		if res[i] > t.max {
			res[i] = t.max
		}
	}
	return &Bid{
		BidType: truestr,
		Bids:    res,
	}
}
