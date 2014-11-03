package auction

type Auction struct {
	Items    int    `json:"items"`
	MaxPrice int    `json:"max_price"`
	Bids     []*Bid `json:"bids"`
}

type Bid struct {
	BidType int   `json:"type"`
	Bids    []int `json:"bids"`
}

func New(items, price int) *Auction {
	return &Auction{
		Items:    items,
		MaxPrice: price,
		Bids:     make([]*Bid, 0),
	}
}

func (a *Auction) Generate(b []Bidder) {
	a.Bids = make([]*Bid, len(b))
	for i, bidder := range b {
		a.Bids[i] = bidder.Bid(a.MaxPrice)
	}
}

func (a *Auction) ClearingPrice() int {
	for i := 0; i < a.MaxPrice; i++ {
		sum := 0
		for _, bids := range a.Bids {
			sum += bids.Bids[i]
		}
		if sum < a.Items {
			return i
		}
	}
	return a.MaxPrice - 1
}

func (a *Auction) BidderPrice(n int) int {
	for i := 0; i < a.MaxPrice; i++ {
		sum := 0
		for j := 0; j < len(a.Bids); j++ {
			if j == n {
				continue
			}
			sum += a.Bids[j].Bids[i]
		}
		if sum < a.Items {
			return i
		}
	}
	return a.MaxPrice - 1
}

//Row returns the winning row results by bids index
func (a *Auction) Row(n int) []int {
	l := len(a.Bids)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = a.Bids[i].Bids[n]
	}
	return res
}

func (a *Auction) Distribute(r int) []int {
	items := a.Items
	row := a.Row(r)
	l := len(row)
	res := make([]int, l)

	for i := 0; i < l && items > 0; i++ {
		max := -1
		for j := 0; j < l; j++ {
			if res[j] == 0 && (max == -1 || row[j] > row[max]) {
				max = j
			}
		}
		if items >= row[max] {
			res[max] = row[max]
			items -= row[max]
		} else {
			res[max] = items
			break
		}
	}

	return res
}
