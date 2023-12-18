package examples

import (
	"fmt"
	"log"

	"github.com/AlexJarrah/Product-Scraper/modules/goat"
)

func goat_() {
	resp, err := goat.FetchGoatProducts("air yeezy", "")
	if err != nil {
		log.Panic(err)
	}

	if len(resp) == 0 {
		log.Panic("no products found")
	}

	fmt.Println("Results:", len(resp))
	for _, p := range resp {
		fmt.Println("Name:", p.Value)
		fmt.Println("SKU:", p.Data.Sku)
		fmt.Printf("Price: $%.2f\n", (float32(p.Data.LowestPriceCents) / 100))
	}
}
