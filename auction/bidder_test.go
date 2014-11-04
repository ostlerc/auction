package auction

import "testing"

func TestCustom(t *testing.T) {
	b := &CustomBidder{money: 20, max: 10}
	if v := b.Bid(5); !ArEq(v.Bids, []int{10, 10, 6, 5, 4}) {
		t.Fatal("Incorrect true bid", v)
	}
}
