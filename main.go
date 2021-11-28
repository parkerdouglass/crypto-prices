package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type MessariResponse struct {
	Data struct {
		MarketData struct {
			PriceUsd float64 `json:"price_usd"`
		} `json:"market_data"`
	} `json:"data"`
}

func main() {
	// Get command-line flags
	var cryptoName string
	flag.StringVar(&cryptoName, "c", "btc", "The cryptocurrency you want the price of")
	flag.Parse()
	// Send API request
	res, err := http.Get("https://data.messari.io/api/v1/assets/" + cryptoName + "/metrics")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// Check API status code
	if res.StatusCode == 404 {
		fmt.Printf("Unsupported cryptocurrency: %s", cryptoName)
		return
	} else if res.StatusCode < 200 || res.StatusCode > 299 {
		panic(fmt.Errorf("API returned status code: %d", res.StatusCode))
	}
	// Parse the JSON reponse
	var messariRes MessariResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&messariRes)
	if err != nil {
		panic(err)
	}
	// Print out the price of the specified crypto
	fmt.Printf("%s price is $%0.10f\n", cryptoName, messariRes.Data.MarketData.PriceUsd)

}
