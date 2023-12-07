package examples

import (
	"fmt"
	"log"

	"github.com/AlexJarrah/Product-Scraper/internal/modules/solebox"
)

func solebox_() {
	resp, err := solebox.FetchSoleboxProducts("black shoes", "")
	if err != nil {
		log.Panic(err)
	}

	for _, product := range resp {
		fmt.Println(product.Name)
		fmt.Printf("$%s", product.Price)
	}
}
