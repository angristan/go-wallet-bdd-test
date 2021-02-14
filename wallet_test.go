package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

func getAssetCategory(categoryString string) (AssetCategory, error) {
	var category AssetCategory
	switch categoryString {
	case "stock":
		category = Stock
	case "crypto":
		category = CryptoCurrency

	case "FIAT":
		category = FIAT
	default:
		return -1, errors.New("invalid category")
	}

	return category, nil
}

func iHaveTheseAssets(myAssets *messages.PickleStepArgument_PickleTable) error {
	assets = []asset{}
	for k, row := range myAssets.Rows {
		// Skip column names
		if k == 0 {
			continue
		}

		// 1/3: Category
		category, err := getAssetCategory(row.Cells[0].Value)
		if err != nil {
			return err
		}

		// 2/3: Ticker
		ticker := row.Cells[1].Value

		// 3/3: Amount
		amount, err := strconv.ParseFloat(row.Cells[2].Value, 64)
		if err != nil {
			return err
		}

		newAsset := asset{
			category: category,
			ticker:   ticker,
			amount:   amount,
		}

		addAsset(newAsset)
	}
	return nil
}

func setAssetPriceUSD(assetName string, assetValue float64) error {
	setAssetPrice(assetName, assetValue)

	return nil
}

func theTotalValueShouldBe(exectedTotalValue float64, currency string) error {
	walletValue, err := getWalletValue(currency)
	if err != nil {
		return err
	}

	if walletValue != exectedTotalValue {
		return errors.New("didn't get exacted wallet value")
	}

	return nil
}

func currencyRateIs(currency string, rate float64) error {
	rates[currency] = rate
	return nil
}

func theTotalValueShouldNotBeUSD(expectedValue float64) error {
	walletValue, err := getWalletValue("USD")
	if err != nil {
		return err
	}
	if walletValue == expectedValue {
		return errors.New("didn't get exacted wallet value: " + fmt.Sprintf("%f", walletValue))
	}
	return nil
}

func theTotalValueShouldNotBeCurrency(expectedValue float64, currency string) error {
	walletValue, err := getWalletValue(currency)
	if err != nil {
		return err
	}
	if walletValue == expectedValue {
		return errors.New("didn't get exacted wallet value: " + fmt.Sprintf("%f", walletValue))
	}
	return nil
}

func theTotalValueInShouldShowAnError(currency string) error {
	_, err := getWalletValue(currency)
	if err == nil {
		return errors.New("expected an error but received none")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	prices = make(map[string]float64)
	rates = make(map[string]float64)
	ctx.Step(`^I have these assets:$`, iHaveTheseAssets)
	ctx.Step(`^"([^"]*)" is worth (\d+) USD$`, setAssetPriceUSD)
	ctx.Step(`^"([^"]*)" rate is (\-*\d+\.\d+)$`, currencyRateIs)
	ctx.Step(`^the total value should be (\d+) ([^"]*)$`, theTotalValueShouldBe)
	ctx.Step(`^the total value should be (\d+) "([^"]*)"$`, theTotalValueShouldBe)
	ctx.Step(`^the total value should not be (\d+)$`, theTotalValueShouldNotBeUSD)
	ctx.Step(`^the total value should not be (\d+) ([^"]*)$`, theTotalValueShouldNotBeCurrency)
	ctx.Step(`^the total value in "([^"]*)" should show an error$`, theTotalValueInShouldShowAnError)

}
