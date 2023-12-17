package examples

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/modules/snipes"
)

func snipes_() {
	resp, err := snipes.FetchSnipesProducts("Dunk", "")
	if err != nil {
		panic(err)
	}

	for _, product := range resp {
		fmt.Printf("%s %s\n", product.Name, product.Color)
		fmt.Println(product.InStock)
		fmt.Println(product.Price.USD)
	}
}
