package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

type Profile struct {
	Price float32 `json:"price"`
}

type Company struct {
	Symbol  string `json:"symbol"`
	Profile Profile
}

var client = &http.Client{Timeout: 10 * time.Second}

func fetch(url string, target interface{}) error {
	r, _ := client.Get(url)
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	runtime.GOMAXPROCS(2)

	tickers := []string{
		"GOOGL",
		"MSFT",
		"AAPL",
		"TSLA",
		"AMZN",
	}

	start := time.Now()
	url := "https://financialmodelingprep.com/api/v3/company/profile/"
	counter := 0

	for _, t := range tickers {
		go func(t string) {
			var c Company
			_ = fetch(url+t, &c)
			fmt.Printf("Company %v with Price %v\n", c.Symbol, c.Profile.Price)
			counter++
		}(t)
	}

	for counter < len(tickers) {
		time.Sleep(10 * time.Millisecond)
	}

	end := time.Since(start)
	fmt.Printf("Execution Time: %s", end)
}
