package main

import (
	"fmt"
	"lld/biddingPlatform/service"
)

func main() {

	am := service.NewAuctionMediator()

	b1 := service.NewBidder(am, "kamlesh")
	b2 := service.NewBidder(am, "tikesh")

	b1.PlaceBid(100)
	b2.PlaceBid(200)

	fmt.Println(am.GetHighestBid())
}
