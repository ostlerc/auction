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
	low, high int
	t         int
}

type TrueBidder struct {
	money int
	max   int
}

func SlowBidder(n int) Bidder {
	return &RangeBidder{
		low:  0,
		high: n / 4,
		t:    Slow,
	}
}

func FastBidder(n int) Bidder {
	return &RangeBidder{
		low:  n / 4,
		high: n / 2,
		t:    Fast,
	}
}

func RandomBidder(n int) Bidder {
	return &RangeBidder{
		low:  0,
		high: n,
		t:    Random,
	}
}

func (r *RangeBidder) Bid(n int) *Bid {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = R.Intn(r.high-r.low) + R.Intn(r.low)
	}
	return &Bid{
		BidType: r.t,
		Bids:    res,
	}
}

func (t *TrueBidder) Bid(n int) *Bid {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = t.money / (i + 1)
		if res[i] > t.max {
			res[i] = t.max
		}
	}
	return &Bid{
		BidType: True,
		Bids:    res,
	}
}

func TypeString(t int) string {
	switch t {
	case Slow:
		return slowstr
	case Fast:
		return faststr
	case Random:
		return randomstr
	case True:
		return truestr
	default:
		panic("Invalid type")
	}
}
