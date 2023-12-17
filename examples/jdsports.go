package examples

import (
	"fmt"
	"log"

	"github.com/AlexJarrah/Product-Scraper/modules/jdsports"
)

func jdsports_() {
	resp, err := jdsports.FetchJDSportsProduct("FN6622-201", "")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(resp.Name)
	fmt.Printf("$%f\n", resp.Price)
}
