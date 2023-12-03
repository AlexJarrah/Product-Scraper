package main

import (
	"fmt"

	products "github.com/quo0001/Product-Scraper"
)

func bodega() {
	link := "https://bdgastore.com/products/palermo-palermo-f-c-39724501"
	product, err := products.Shopify(link, "")
	if err != nil {
		panic(err)
	}

	fmt.Println("Product Information for", product.Product.Title)
	for _, p := range product.Product.Variants {
		fmt.Println("Variant:", p.ID)
		fmt.Println("Title:", p.Title)
		fmt.Println("Stock:", p.InventoryQuantity)
	}
}

func newburycomics() {
	link := "https://www.newburycomics.com/collections/newbury-comics-merch/products/newbury_comics-newbury_comics_toothface_logo_tie-dye_t-shirt?variant=42267798634676"
	product, err := products.Shopify(link, "")
	if err != nil {
		panic(err)
	}

	fmt.Println("Product Information for", product.Product.Title)
	for _, p := range product.Product.Variants {
		fmt.Println("Variant:", p.ID)
		fmt.Println("Title:", p.Title)
		fmt.Println("Stock:", p.InventoryQuantity)
	}
}

func eminem() {
	link := "https://shop.eminem.com/products/eminem-x-fortnite-radio-vinyl"
	product, err := products.Shopify(link, "")
	if err != nil {
		panic(err)
	}

	fmt.Println("Product Information for", product.Product.Title)
	for _, p := range product.Product.Variants {
		fmt.Println("Variant:", p.ID)
		fmt.Println("Title:", p.Title)
		fmt.Println("Stock:", p.InventoryQuantity)
	}
}
