package main

import (
	"fmt"
	"log"

	gecko "github.com/vladivolo/go-gecko/v3"
)

func main() {
	tickers, err := gecko.CoinsIDTickers("bitcoin", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tickers.Tickers)
	tickers, err = gecko.CoinsIDTickers("bitcoin", 1)
	fmt.Println(len(tickers.Tickers))
}
