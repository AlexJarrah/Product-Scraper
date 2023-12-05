package examples

import (
	"fmt"

	"github.com/AlexJarrah/Product-Scraper/internal/modules/nike"
)

func nike_() {
	resp, err := nike.FetchNikeProducts("FN6622_201", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp[0].Title)
	fmt.Println(resp[0].InStock)
	fmt.Println(resp[0].Price.CurrentPrice)
}
