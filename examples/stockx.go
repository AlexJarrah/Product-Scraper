package examples

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/modules/stockx"
)

func stockx_() {
	resp, err := stockx.FetchStockXProducts("508214-660", "")
	if err != nil {
		panic(err)
	}

	for _, p := range resp {
		fmt.Println("Name:", p.Node.Name)
		fmt.Println("Last Sale:", p.Node.Market.SalesInformation.LastSale)
		fmt.Println("Lowest Ask:", p.Node.Market.BidAskData.LowestAsk)
		fmt.Println("Highest Bid:", p.Node.Market.BidAskData.HighestBid)
	}
}
