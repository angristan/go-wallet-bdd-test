package main

import "github.com/piquette/finance-go/quote"

func setAssetPrice(ticker string, price float64) {
	prices[ticker] = price
}

func getAssetPrice(a asset) (float64, error) {
	var ticker string

	if a.category == CryptoCurrency {
		ticker = a.ticker + "-USD"
	} else {
		ticker = a.ticker
	}

	if prices[ticker] == 0 {
		q, err := quote.Get(ticker)
		if err != nil {
			return 0, err
		}
		setAssetPrice(ticker, q.RegularMarketPrice)
	}

	return prices[ticker], nil
}
