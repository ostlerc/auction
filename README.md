auction
=======
multi-unit vickrey auction analysis software

agen
====
  View single auction results.

    Usage of agen:
    -f=0: Number of fast bidders in the auction.
    -items=30: Number of items in auction
    -o="csv": Output type [json,csv]
    -price=10: Maximum price of auction
    -r=3: Number of random bidders in the auction.
    -s=0: Number of slow bidders in the auction.
    -t=0: Number of tru bidders in the auction.
    -tm=100: Maximum amount of money a true bidder can auction.

  sample csv output:

     Items        30
     Clear Price  10
     Price        r1  r2  r3
     1            29  29  25
     2            28  28  19
     3            27  28  18
     4            26  28  15
     5            25  28  14
     6            25  27  8
     7            24  26  7
     8            24  26  4
     9            23  26  1
     10           22  25  0
     Items won    4   26  0
     Cost         9   8   0
     Utility      1   2   0

 Note:  Header values represent bidder type and number types are as follows:
   'r' for random, 'f' for fast, 's' for slow, 't' for true (original).

recap
=====
  simulate and recap several auctions

    Usage of recap:
    -b="5,13,21": csv of bidder rows to output
    -items=30: Number of items in auction
    -n=5: Auctions to aggregate
    -price=10: Maximum price of auction
    -tm=100: Maximum amount of money a true bidder can auction.

  sample csv output:

     Bidders  rand(Ct $ Util)           fast(Ct $ Util)           slow(Ct $ Util)           all(rand/fast/slow/orig)
     5        0.6 1.4 0.4(4.9 2.1 0.4)  2.0 1.0 0.2(4.7 2.4 0.4)  2.2 1.8 0.2(4.6 4.7 0.3)  1.2 2.8 0.2(6.0 1.6 0.4 / 6.0 4.0 0.0 / 5.4 3.2 1.0 / 5.4 2.8 0.4)
     13       0.0 0.0 0.0(1.9 0.6 0.0)  1.0 1.4 0.0(1.9 1.2 0.0)  0.0 0.0 0.0(1.9 1.7 0.0)  0.0 0.0 0.0(2.3 1.5 0.1 / 2.3 1.5 0.0 / 2.3 2.7 0.1 / 2.3 0.8 0.0)
     21       0.0 0.0 0.0(1.2 0.4 0.0)  0.6 3.2 0.0(1.2 1.2 0.1)  0.0 0.0 0.0(1.2 0.9 0.0)  0.0 0.0 0.0(1.4 1.0 0.0 / 1.4 1.0 0.0 / 1.4 0.5 0.0 / 1.4 1.0 0.0)

 Note: this csv output has been pretty printed in column format

report
======

My original bidding strategy involves a maximum amount of cash and ratio of items / cost a bidder is willing to accept.
They always bid truthfully. An example output for my report can be found in the 'recap' section of this README. After
running recap a few times, it is apparent that my strategy does not win very often. When it does win, the utility is very low.
When there are few bidders the possibility of getting a better utility increases. In general, when against fast he will do fairly well.
When up againt just random bidders it is ... hit and miss as random would suggest. The slow bidders are hard to get very high Utility
against. When doing a mix the probability is higher than if just against fast bidders. But the overall utility gained is dimished as
other strategies are combined.
