package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/ostlerc/auction/auction"
)

var (
	items = flag.Int("items", 30, "Number of items in auction")
	price = flag.Int("price", 10, "Maximum price of auction")

	random = flag.Int("random", 1, "Number of random bidders in the auction.")
	fast   = flag.Int("fast", 1, "Number of fast bidders in the auction.")
	slow   = flag.Int("slow", 1, "Number of slow bidders in the auction.")
	tru    = flag.Int("tru", 1, "Number of tru bidders in the auction.")
	tmm    = flag.Int("tru_max_money", *price*10, "Maximum amount of money a true bidder auctions.")
)

func bidders() []auction.Bidder {
	res := make([]auction.Bidder, 0)
	for i := 0; i < *random; i++ {
		res = append(res, auction.RandomBidder(*items))
	}
	for i := 0; i < *slow; i++ {
		res = append(res, auction.SlowBidder(*items))
	}
	for i := 0; i < *fast; i++ {
		res = append(res, auction.FastBidder(*items))
	}
	for i := 0; i < *tru; i++ {
		res = append(res, auction.TrueBidder(*tmm, *items))
	}
	return res
}

func main() {
	flag.Parse()
	a := auction.New(*items, *price)
	a.Generate(bidders())
	dat, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dat))
}
