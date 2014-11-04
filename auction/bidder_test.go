package auction

import "testing"

func TestCustom(t *testing.T) {
	b := &CustomBidder{money: 20, max: 2, i: 10}
	if v := b.Bid(5); !ArEq(v.Bids, []int{10, 5, 3, 2, 2}) {
		t.Fatal("Incorrect true bid", v)
	}
}
