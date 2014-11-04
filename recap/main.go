package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ostlerc/auction/auction"
)

var (
	items = flag.Int("items", 30, "Number of items in auction")
	price = flag.Int("price", 10, "Maximum price of auction")
	tmm   = flag.Int("tm", *price*10, "Maximum amount of money a true bidder can auction.")

	b       = flag.String("b", "5,13,21", "csv of bidder rows to output")
	n       = flag.Int("n", 5, "Auctions to aggregate")
	bidders = []genBidder{randomBidders, fastBidders, slowBidders, allBidders}
)

type genBidder func(int) []auction.Bidder

func randomBidders(n int) []auction.Bidder {
	bidders := make([]auction.Bidder, 0)
	for i := 0; i < n; i++ {
		bidders = append(bidders, auction.RandomBidder(*items))
	}
	return bidders
}

func fastBidders(n int) []auction.Bidder {
	bidders := make([]auction.Bidder, 0)
	for i := 0; i < n; i++ {
		bidders = append(bidders, auction.FastBidder(*items))
	}
	return bidders
}

func slowBidders(n int) []auction.Bidder {
	bidders := make([]auction.Bidder, 0)
	for i := 0; i < n; i++ {
		bidders = append(bidders, auction.SlowBidder(*items))
	}
	return bidders
}

func origBidders(n int) []auction.Bidder {
	bidders := make([]auction.Bidder, 0)
	for i := 0; i < n; i++ {
		bidders = append(bidders, auction.TrueBidder(*tmm, *items, *price))
	}
	return bidders
}

func allBidders(n int) []auction.Bidder {
	bidders := make([]auction.Bidder, 0)
	bidders = append(bidders, randomBidders(n/4)...)
	n++
	bidders = append(bidders, fastBidders(n/4)...)
	n++
	bidders = append(bidders, slowBidders(n/4)...)
	n++
	bidders = append(bidders, origBidders(n/4)...)
	return bidders
}

func resultsStr(r []*auction.Result) string {
	price := make(auction.Ints, len(r[0].Prices))
	ct := make(auction.Ints, len(r[0].Distribution))
	util := make(auction.Ints, len(r[0].Utility))

	for _, v := range r {
		price[0] += v.Prices[0]
		ct[0] += v.Distribution[0]
		util[0] += v.Utility[0]
		price = append(price, v.Prices[1:]...)
		ct = append(ct, v.Distribution[1:]...)
		util = append(util, v.Utility[1:]...)
	}
	fprice := float64(price[0]) / float64(len(r))
	fct := float64(ct[0]) / float64(len(r))
	futil := float64(util[0]) / float64(len(r))
	return fmt.Sprintf("%.1f %.1f %.1f(%.1f %.1f %.1f),",
		fct,
		fprice,
		futil,
		ct.Ave(0),
		price.Ave(0),
		util.Ave(0))
}

func allResultsStr(r []*auction.Result) string {
	var (
		//results[bid group][price/ct/util]
		results            [4][3]auction.Ints
		fprice, fct, futil float64
	)

	at := 0
	for i := 0; i < 4; i++ { //four groups
		results[i][0] = make(auction.Ints, 0)
		results[i][1] = make(auction.Ints, 0)
		results[i][2] = make(auction.Ints, 0)
		for j := 0; j < (len(r)+i-1)/4; j++ { //find distributed group range
			fprice += float64(r[at].Prices[0])
			fct += float64(r[at].Distribution[0])
			futil += float64(r[at].Utility[0])
			results[i][0] = append(results[i][0], r[at].Distribution[1:]...)
			results[i][1] = append(results[i][1], r[at].Prices[1:]...)
			results[i][2] = append(results[i][2], r[at].Utility[1:]...)
			at++
		}
	}

	fprice /= float64(len(r))
	fct /= float64(len(r))
	futil /= float64(len(r))

	res := ""
	res += fmt.Sprintf("%.1f %.1f %.1f(",
		fct,
		fprice,
		futil)

	for i := 0; i < 4; i++ {
		res += fmt.Sprintf("%.1f %.1f %.1f",
			results[i][0].Ave(-1),
			results[i][1].Ave(-1),
			results[i][2].Ave(-1))
		if i != 3 {
			res += " / "
		}
	}
	res += ")"

	return res
}

func main() {
	flag.Parse()
	fmt.Println("Bidders,rand(Ct $ Util),fast(Ct $ Util),slow(Ct $ Util),all(rand/fast/slow/orig)")
	a := auction.New(*items, *price)
	for _, v := range strings.Split(*b, ",") {
		fmt.Print(v, ",")
		for i, b := range bidders {
			results := make([]*auction.Result, 0)
			for x := 0; x < *n; x++ {
				bidders := make([]auction.Bidder, 1)
				oBid := auction.TrueBidder(*tmm, *items, *price)
				bidders[0] = oBid
				n, _ := strconv.Atoi(v)
				bidders = append(bidders, b(n)...)
				a.Conduct(bidders)
				results = append(results, a.Result())
			}
			if i+1 != len(bidders) {
				fmt.Print(resultsStr(results))
			} else {
				fmt.Print(allResultsStr(results))
			}
		}
		fmt.Println("")
	}
}
