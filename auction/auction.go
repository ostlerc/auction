package auction

import (
	"fmt"
	"strconv"
	"strings"
)

// Bids encapsulates a list of bids, some helper functions apply
type Bids []*Bid

// Auction contains all information for an auction
type Auction struct {
	Items    int     `json:"items"`
	MaxPrice int     `json:"max_price"`
	Bids     Bids    `json:"bids"`
	Res      *Result `json:"result"`
}

// Bid contains descending order of bids for one bidder
type Bid struct {
	Type string `json:"type"`
	Bids Ints   `json:"bids"`
}

// Result contains all auction result information
type Result struct {
	Distribution Ints `json:"distribution"`
	Prices       Ints `json:"prices"`
	Utility      Ints `json:"utility"`
}

// New creates a new auction with given item and price
func New(items, price int) *Auction {
	return &Auction{
		Items:    items,
		MaxPrice: price,
		Bids:     make(Bids, 0),
	}
}

// Conduct conducts an auction based on bidders
func (a *Auction) Conduct(b []Bidder) {
	a.Bids = make(Bids, len(b))
	for i, bidder := range b {
		a.Bids[i] = bidder.Bid(a.MaxPrice)
	}
	a.Res = a.Result()
}

// ClearingPrice returns the price where there is no longer enough
// demand for the object to sell all objects
func (a *Auction) ClearingPrice() int {
	for i := 0; i < a.MaxPrice; i++ {
		if a.Row(i).Sum() < a.Items {
			return i
		}
	}
	return a.MaxPrice - 1
}

// BidderPrice returns the price for bidder n
func (a *Auction) BidderPrice(n int) int {
	for i := 0; i < a.MaxPrice; i++ {
		sum := a.Row(i).Sum()
		if sum-a.Bids[n].Bids[i] < a.Items {
			return i + 1
		}
	}
	return a.MaxPrice
}

// BidderPrices returns an array of prices each bidder needs to pay
func (a *Auction) BidderPrices() Ints {
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
		Utility:      a.Utility(),
	}
}

//Row returns the winning row results by bids index
func (a *Auction) Row(n int) Ints {
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

// Utility returns list of utility per bidder
func (a *Auction) Utility() Ints {
	res := make(Ints, len(a.Bids))
	cp := a.ClearingPrice() + 1
	for i, v := range a.BidderPrices() {
		res[i] = cp - v
	}
	return res
}

// CSV returns auction in CSV form
func (a *Auction) CSV() string {
	res := ""
	res += fmt.Sprintf("Items,%v\n", a.Items)
	res += fmt.Sprintf("Clear Price,%v\n", a.ClearingPrice()+1)
	res += "Price," + strings.Join(a.Bids.Types(), ",") + "\n"
	for i := 0; i < a.MaxPrice; i++ {
		res += strconv.Itoa(i+1) + "," +
			strings.Join(a.Row(i).String(), ",") + "\n"
	}
	res += a.Res.CSV()
	return res
}

// CSV returns a result in CSV form
func (r *Result) CSV() string {
	res := "Items won," + strings.Join(r.Distribution.String(), ",") + "\n"
	res += "Cost," + strings.Join(r.Prices.String(), ",") + "\n"
	res += "Utility," + strings.Join(r.Utility.String(), ",")
	return res
}

// Types returns an array of types from Bids
func (b Bids) Types() []string {
	res := make([]string, len(b))
	for i, v := range b {
		res[i] = string(v.Type[0]) + strconv.Itoa(i+1)
	}
	return res
}
