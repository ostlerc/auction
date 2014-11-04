package main

import "github.com/ostlerc/auction/auction"

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
