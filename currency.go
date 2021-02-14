package main

import (
	"errors"
	"github.com/asvvvad/exchange"
)

func isCurrencyValid(currency string) (bool, error) {
	ex := exchange.New("")
	validCurrencies, err := ex.LatestRatesAll()
	if err != nil {
		return false, err
	}
	for c, _ := range validCurrencies {
		if c == currency {
			return true, nil
		}
	}
	return false, nil
}
func getExchangeRate(currency string) (float64, error) {
	valid, err := isCurrencyValid(currency)
	if err != nil {
		return 0, err
	}
	if !valid {
		return 0, errors.New("invalid currency")
	}
	ex := exchange.New("USD")
	rate, err := ex.ConvertTo(currency, 1)
	if err != nil {
		return 0, err
	}
	rateFloat, _ := rate.Float64()

	return rateFloat, err
}

func convertUSDto(amount float64, currency string) (float64, error) {
	if rates[currency] == 0 {
		exchangeRate, err := getExchangeRate(currency)
		if err != nil {
			return 0, err
		}
		rates[currency] = exchangeRate
	}
	return amount * rates[currency], nil
}
