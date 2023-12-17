package examples

import (
	"fmt"
	"log"

	"github.com/AlexJarrah/Product-Scraper/modules/supreme"
)

func supremeCollections() {
	collections, err := supreme.FetchSupremeCollections("")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Products:", collections.AllProductsCount)
	for _, product := range collections.Products {
		fmt.Println("Name:", product.Title)
		fmt.Println("In Stock:", product.Available)
		fmt.Printf("Price: $%.2f\n", (float32(product.Price) / 100))
	}
}

func supremeSeason() {
	szn, err := supreme.FetchSupremeSeason("fallwinter2023", "")
	if err != nil {
		log.Panic(err)
	}

	for _, group := range szn.Preview.Groups {
		for _, product := range group.Products {
			fmt.Println(product.Title)
		}
	}
}
