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

	random = flag.Int("r", 3, "Number of random bidders in the auction.")
	fast   = flag.Int("f", 0, "Number of fast bidders in the auction.")
	slow   = flag.Int("s", 0, "Number of slow bidders in the auction.")
	tru    = flag.Int("t", 0, "Number of tru bidders in the auction.")
	tmm    = flag.Int("tm", *price*10, "Maximum amount of money a true bidder can auction.")

	o = flag.String("o", "csv", "Output type [json,csv]")
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
		res = append(res, auction.TrueBidder(*tmm, *items, *price))
	}
	return res
}

func main() {
	flag.Parse()
	if *o != "csv" && *o != "json" {
		log.Fatal("invalid output type '", *o, "'")
	}

	a := auction.New(*items, *price)
	a.Conduct(bidders())

	if *o == "json" {
		dat, err := json.Marshal(a)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(dat))
	} else { //csv
		fmt.Println(a.CSV())
	}
}
