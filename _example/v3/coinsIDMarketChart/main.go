package main

import (
	"fmt"
	"log"
	"time"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	btc_usd, err := gecko.CoinsIDMarketChart("bitcoin", "usd", "365")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TotalVolumes: %d\n", len(*btc_usd.TotalVolumes))
	for _, v := range *btc_usd.TotalVolumes {
		fmt.Printf("%s:%.04f\n", time.Unix(int64(v[0])/1000, int64(v[0])%1000).UTC().Format(time.RFC3339), v[1])
	}

	/*
		btc_eur, err := gecko.CoinsIDMarketChart("bitcoin", "eur", "365")
		if err != nil {
			log.Fatal(err)
		}

		orig := [][2]float32{}

		for _, v := range *btc_usd.Prices {
			orig = append(orig, [2]float32{v[0], 1 / v[1]})
		}

		for _, v := range orig {
			fmt.Printf("%s:%.08f\n", time.Unix(int64(v[0])/1000, int64(v[0])%1000).UTC().Format(time.RFC3339), v[1])
		}

		fmt.Printf("----------------------------\n")

		for idx, v := range *btc_eur.Prices {
			fmt.Printf("%s:%.08f\n", time.Unix(int64(v[0])/1000, int64(v[0])%1000).UTC().Format(time.RFC3339), v[1]*orig[idx][1])
		}
	*/
	/*

			fmt.Printf("Prices: %d\n", len(*m.Prices))
			for _, v := range *m.Prices {
				fmt.Printf("%s:%.04f\n", time.Unix(int64(v[0])/1000, int64(v[0])%1000).UTC().Format(time.RFC3339), v[1])
			}
		    fmt.Printf("MarketCaps: %d\n", len(*m.MarketCaps))
		    for _, v := range *m.MarketCaps {
		        fmt.Printf("%s:%.04f\n", time.Unix(int64(v[0])/1000, int64(v[0])%1000).UTC().Format(time.RFC3339), v[1])
		    }

		    fmt.Printf("TotalVolumes: %d\n", len(*m.TotalVolumes))
		    for _, v := range *m.TotalVolumes {
		        fmt.Printf("%s:%.04f\n", time.Unix(int64(v[0])/1000, int64(v[0])%1000).UTC().Format(time.RFC3339), v[1])
		    }
	*/
}
