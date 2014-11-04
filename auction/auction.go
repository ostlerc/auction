package auction

// Auction contains all information for an auction
type Auction struct {
	Items    int    `json:"items"`
	MaxPrice int    `json:"max_price"`
	Bids     []*Bid `json:"bids"`
}

// Bid contains descending order of bids for one bidder
type Bid struct {
	BidType string `json:"type"`
	Bids    []int  `json:"bids"`
}

// Result contains all auction result information
type Result struct {
	Distribution []int `json:"distribution"`
	Prices       []int `json:"prices"`
}

// New creates a new auction with given item and price
func New(items, price int) *Auction {
	return &Auction{
		Items:    items,
		MaxPrice: price,
		Bids:     make([]*Bid, 0),
	}
}

// Generate sets bidders for an auction
func (a *Auction) Generate(b []Bidder) {
	a.Bids = make([]*Bid, len(b))
	for i, bidder := range b {
		a.Bids[i] = bidder.Bid(a.MaxPrice)
	}
}

// ClearingPrice returns the price where there is no longer enough
// demand for the object to sell all objects
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

// BidderPrice returns the price for bidder n
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
			return a.MaxPrice - i
		}
	}
	return a.MaxPrice - 1
}

// BidderPrices returns an array of prices each bidder needs to pay
func (a *Auction) BidderPrices() []int {
	res := make([]int, len(a.Bids))
	for i := 0; i < len(a.Bids); i++ {
		res[i] = a.BidderPrice(i)
	}
	return res
}

// Result returns a result struct of the auction
func (a *Auction) Result() *Result {
	return &Result{
		Distribution: a.Distribute(a.ClearingPrice() - 1),
		Prices:       a.BidderPrices(),
	}
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

// Distribute returns an array of how items are distributed in the auction
func (a *Auction) Distribute(r int) []int {
	if r == -1 {
		r = 0
	}
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
