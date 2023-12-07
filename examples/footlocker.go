package examples

import (
	"fmt"
	"log"

	"github.com/AlexJarrah/Product-Scraper/internal/modules/footlocker"
)

func footlocker_() {
	resp, err := footlocker.FetchFootlockerProduct("W2288111", "")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(resp.Model.Name)
	fmt.Println(resp.Style.Price.CurrencyIso)
}
