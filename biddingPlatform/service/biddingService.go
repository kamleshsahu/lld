package service

import "fmt"

type IBidder interface {
	PlaceBid(amount int)
	OnNotify(bidderId string, amount int)
	GetId() string
}

type Bidder struct {
	Id              string
	Name            string
	AuctionMediator IAuctionMediator
}

func (b Bidder) GetId() string {
	return b.Name
}

func (b Bidder) PlaceBid(amount int) {
	b.AuctionMediator.PlaceBid(b.Name, amount)
}

func (b Bidder) OnNotify(bidderId string, amount int) {
	fmt.Println("Bidder ", b.Name, fmt.Sprintf(" got notified with %s:%d ", bidderId, amount))
}

func NewBidder(mediator IAuctionMediator, name string) IBidder {
	bidder := &Bidder{AuctionMediator: mediator, Name: name, Id: name}
	mediator.AddBidder(bidder)
	return bidder
}

type IAuctionMediator interface {
	AddBidder(b IBidder)
	PlaceBid(bidderId string, amount int)
	GetHighestBid() Bid
}

type Bid struct {
	bidderId string
	amount   int
}

type AuctionMediator struct {
	bidders    []IBidder
	bids       []Bid
	highestBid Bid
}

func (a *AuctionMediator) GetHighestBid() Bid {
	return a.highestBid
}

func (a *AuctionMediator) AddBidder(b IBidder) {
	a.bidders = append(a.bidders, b)
}

func (a *AuctionMediator) PlaceBid(bidderId string, amount int) {
	if amount < a.highestBid.amount {
		fmt.Println("bid not allowed")
		return
	}
	bid := Bid{bidderId: bidderId, amount: amount}
	a.bids = append(a.bids, bid)
	a.highestBid = bid
	fmt.Println("bid placed by ", bidderId, " with amount ", amount)
	for _, bidder := range a.bidders {
		if bidderId == bidder.(IBidder).GetId() {
			continue
		}
		bidder.OnNotify(bidderId, amount)
	}
}

func NewAuctionMediator() IAuctionMediator {
	return &AuctionMediator{bidders: make([]IBidder, 0)}
}
