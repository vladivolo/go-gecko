package main

import (
	"fmt"
	"log"

	gecko "github.com/vladivolo/go-gecko/v3"
)

var avail = map[string]bool{"btc": true, "usd": true, "eur": true, "rub": true, "aed": true, "gbp": true, "cny": true}
var updown = 0.1

func main() {
	coin, err := gecko.CoinsID("bitcoin", true, true, true, true, true, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%V\n", coin.MarketData.CurrentPrice)
	return

	// PRICE CHANGE PERCENTILE
	zero_level := float64(0)

	fmt.Printf("btc: ")
	for symbol, value := range coin.MarketData.PriceChangePercentage1hInCurrency {
		if _, ok := avail[symbol]; ok {
			if symbol == "usd" {
				zero_level = value
			}
			fmt.Printf("%s:%f", symbol, value)
			fmt.Print("% ")
		}
	}
	fmt.Println()

	// Tkey
	for symbol, value := range coin.MarketData.PriceChangePercentage1hInCurrency {
		if _, ok := avail[symbol]; ok {
			fmt.Printf("%s:%f", symbol, value-zero_level)
			fmt.Print("% ")
		}
	}
	fmt.Println()

	return

	//PRICE
	var (
		mul            float64
		tkey_price_map = map[string]float64{}
	)

	fmt.Printf("btc: ")
	for symbol, value := range coin.MarketData.CurrentPrice {
		if _, ok := avail[symbol]; ok {

			if symbol == "usd" {
				mul = float64(1) / value
			}

			fmt.Printf("%s:%f ", symbol, value)
		}
	}
	fmt.Println()

	fmt.Printf("tkey: ")
	for symbol, value := range coin.MarketData.CurrentPrice {
		if _, ok := avail[symbol]; ok {
			fmt.Printf("%s:%f ", symbol, value*mul)
			tkey_price_map[symbol] = value * mul
		}
	}
	fmt.Println()

	// CHANGE PRICE
	usd_price, _ := coin.MarketData.CurrentPrice["usd"]
	usd_change, _ := coin.MarketData.PriceChange24hInCurrency["usd"]

	fmt.Printf("TKEY: %f %f\n", usd_price, usd_change)

	fmt.Printf("btc change: ")
	for symbol, value := range coin.MarketData.PriceChange24hInCurrency {
		if _, ok := avail[symbol]; ok {
			fmt.Printf("%s:%f ", symbol, value)
		}
	}
	fmt.Println()

	zero_level = usd_change / usd_price

	fmt.Printf("tkey change: ")
	for symbol, change := range coin.MarketData.PriceChange24hInCurrency {
		if _, ok := avail[symbol]; ok {
			price, _ := coin.MarketData.CurrentPrice[symbol]
			tk_s, _ := tkey_price_map[symbol]

			fmt.Printf("%s:%f ", symbol, tk_s*(change/price-zero_level))
		}
	}
	fmt.Println()

	//fmt.Printf("%V\n", coin.MarketData.CurrentPrice)
}
