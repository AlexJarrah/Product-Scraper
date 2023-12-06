package examples

import (
	"fmt"
	"log"

	"github.com/AlexJarrah/Product-Scraper/internal/modules/adidas"
)

func adidas_() {
	skuList := []string{"IG6173", "FY7756"}
	resp, err := adidas.FetchAdidasProducts(skuList, "")
	if err != nil {
		log.Panic(err)
	}

	if len(resp) == 0 {
		log.Panic("no products found")
	} else if len(resp[0].ProductLinkList) == 0 {
		log.Panic("no products found")
	}

	fmt.Println(resp[0].Name)
	fmt.Println(resp[0].ProductLinkList[0].AvailableSkus)
	fmt.Println(resp[0].PricingInformation.CurrentPrice)
}
