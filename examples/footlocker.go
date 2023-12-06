package examples

import (
	"fmt"
	"log"

	"github.com/AlexJarrah/Product-Scraper/internal/modules/footlocker"
)

func footlocker_() {
	resp, err := footlocker.FetchFootlockerProduct("J4210001", "")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(resp.Name)
	fmt.Println(resp.Offers[0].Price)
}
