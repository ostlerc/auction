package auction

import "testing"

func TestClearingPrice(t *testing.T) {
	a := New(30, 5)
	a.Bids = []*Bid{
		&Bid{Bids: []int{6, 5, 4, 3, 2}},
		&Bid{Bids: []int{6, 5, 4, 3, 2}},
		&Bid{Bids: []int{6, 5, 4, 3, 2}},
		&Bid{Bids: []int{6, 5, 4, 3, 2}},
		&Bid{Bids: []int{6, 5, 4, 3, 2}},
		&Bid{Bids: []int{6, 5, 4, 3, 2}},
		&Bid{Bids: []int{6, 5, 4, 3, 2}},
	}

	if p := a.ClearingPrice(); p != 2 {
		t.Fatal(p)
	}

	a.Bids = []*Bid{
		&Bid{Bids: []int{4, 3, 2, 1, 0}},
		&Bid{Bids: []int{4, 3, 2, 1, 0}},
		&Bid{Bids: []int{4, 3, 2, 1, 0}},
		&Bid{Bids: []int{4, 3, 2, 1, 0}},
		&Bid{Bids: []int{4, 3, 2, 1, 0}},
		&Bid{Bids: []int{4, 3, 2, 1, 0}},
		&Bid{Bids: []int{4, 3, 2, 1, 0}},
	}

	if p := a.ClearingPrice(); p != 0 {
		t.Fatal(p)
	}

	a.Bids = []*Bid{
		&Bid{Bids: []int{7, 7, 7, 7, 7}},
		&Bid{Bids: []int{7, 7, 7, 7, 7}},
		&Bid{Bids: []int{7, 7, 7, 7, 7}},
		&Bid{Bids: []int{7, 7, 7, 7, 7}},
		&Bid{Bids: []int{7, 7, 7, 7, 7}},
		&Bid{Bids: []int{7, 7, 7, 7, 7}},
		&Bid{Bids: []int{7, 7, 7, 7, 7}},
	}

	if p := a.ClearingPrice(); p != 4 {
		t.Fatal(p)
	}
}

func TestBidderPrice(t *testing.T) {
	a := New(10, 5)
	a.Bids = []*Bid{
		&Bid{Bids: []int{10, 10, 4, 3, 2}},
		&Bid{Bids: []int{7, 4, 4, 3, 2}},
		&Bid{Bids: []int{4, 2, 1, 1, 1}},
	}

	if p := a.BidderPrice(0); p != 1 {
		t.Fatal(p)
	}
	if p := a.BidderPrice(1); p != 2 {
		t.Fatal(p)
	}
	if p := a.BidderPrice(2); p != 2 {
		t.Fatal(p)
	}
}

func TestRow(t *testing.T) {
	a := New(30, 5)
	a.Bids = []*Bid{
		&Bid{Bids: []int{1, 2, 3, 4, 5}},
		&Bid{Bids: []int{5, 4, 3, 2, 1}},
	}

	if v := a.Row(0); !ArEq([]int{1, 5}, v) {
		t.Fatal(v)
	}

	if v := a.Row(1); !ArEq([]int{2, 4}, v) {
		t.Fatal(v)
	}

	if v := a.Row(2); !ArEq([]int{3, 3}, v) {
		t.Fatal(v)
	}

	if v := a.Row(3); !ArEq([]int{4, 2}, v) {
		t.Fatal(v)
	}

	if v := a.Row(4); !ArEq([]int{5, 1}, v) {
		t.Fatal(v)
	}
}

func TestDistribute(t *testing.T) {
	a := New(5, 5)
	a.Bids = []*Bid{
		&Bid{Bids: []int{1, 2, 3, 4, 5}},
		&Bid{Bids: []int{5, 4, 3, 2, 1}},
	}

	if v := a.Distribute(0); !ArEq([]int{0, 5}, v) {
		t.Fatal(v)
	}
	if v := a.Distribute(1); !ArEq([]int{1, 4}, v) {
		t.Fatal(v)
	}
	if v := a.Distribute(2); !ArEq([]int{3, 2}, v) {
		t.Fatal(v)
	}
	if v := a.Distribute(3); !ArEq([]int{4, 1}, v) {
		t.Fatal(v)
	}
	if v := a.Distribute(4); !ArEq([]int{5, 0}, v) {
		t.Fatal(v)
	}

	a.Items = 4
	if v := a.Distribute(0); !ArEq([]int{0, 4}, v) {
		t.Fatal(v)
	}
	if v := a.Distribute(1); !ArEq([]int{0, 4}, v) {
		t.Fatal(v)
	}
	if v := a.Distribute(2); !ArEq([]int{3, 1}, v) {
		t.Fatal(v)
	}
	if v := a.Distribute(3); !ArEq([]int{4, 0}, v) {
		t.Fatal(v)
	}
	if v := a.Distribute(4); !ArEq([]int{4, 0}, v) {
		t.Fatal(v)
	}
}
