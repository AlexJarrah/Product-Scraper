package examples

import (
	"fmt"
	"log"

	"github.com/AlexJarrah/Product-Scraper/modules/nike"
)

func nike_() {
	resp, err := nike.FetchNikeProducts("FN6622_201", "")
	if err != nil {
		log.Panic(err)
	}

	if len(resp) == 0 {
		log.Panic("no products found")
	}

	fmt.Println(resp[0].Title)
	fmt.Println(resp[0].InStock)
	fmt.Println(resp[0].Price.CurrentPrice)
}
