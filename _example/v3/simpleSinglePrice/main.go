package main

import (
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	price, err := gecko.SimpleSinglePrice("eur", "usd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%s is worth %f %s", price.ID, price.MarketPrice, price.Currency))
}
