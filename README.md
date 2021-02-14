# Simple wallet using Behavior-driven development

## Assignement

> Given a Wallet containing Stocks, build a function that compute the value of wallet in a currency.
>
> The Stocks have a quantity and a StockType. The StockType can be for example petroleum, Euros, bitcoins and Dollars.
>
>To value the portfolio in a Currency you can use external api to provide rate exchanges

## Running tests

This program is using Cucumber for Golang: [Godog](https://github.com/cucumber/godog).

Install godog:

```shell
go get github.com/cucumber/godog/cmd/godog@v0.11.0
```

Run Gherking scenarios:

```shell
godog
```

### Output

```
➜  wallet git:(main) godog
Feature: get wallet value
  Get value of all assets in wallet

  Scenario Outline: Get wallet value in target currency with specified rate # features/get_wallet_value.feature:4
    Given I have these assets:                                              # wallet_test.go:28 -> iHaveTheseAssets
      | type  | ticker | quantity |
      | stock | TSLA   | 1.1      |
      | stock | AAPL   | 10       |
    And "TSLA" is worth 1000 USD                                            # wallet_test.go:62 -> setAssetPriceUSD
    And "AAPL" is worth 130 USD                                             # wallet_test.go:62 -> setAssetPriceUSD
    And "<Currency>" rate is <Rate>                                         # wallet_test.go:81 -> currencyRateIs
    Then the total value should be <Value> "<Currency>"                     # wallet_test.go:68 -> theTotalValueShouldBe

    Examples:
      | Currency | Rate   | Value  |
      | EUR      | 0.83   | 1992   |
      | JPY      | 104.95 | 251880 |
      | USD      | 1.0    | 2400   |

  Scenario: Get wallet value in invalid currency       # features/get_wallet_value.feature:20
    Given I have these assets:                         # wallet_test.go:28 -> iHaveTheseAssets
      | type  | ticker | quantity |
      | stock | TSLA   | 1.1      |
      | stock | AAPL   | 10       |
    Then the total value in "LOL" should show an error # wallet_test.go:108 -> theTotalValueInShouldShowAnError

  Scenario: Get wallet value without assets # features/get_wallet_value.feature:27
    Given I have these assets:              # wallet_test.go:28 -> iHaveTheseAssets
      | type | ticker | quantity |
    Then the total value should be 0 USD    # wallet_test.go:68 -> theTotalValueShouldBe

  Scenario: Get wallet value in USD with stocks values from Yahoo Finance # features/get_wallet_value.feature:32
    Given I have these assets:                                            # wallet_test.go:28 -> iHaveTheseAssets
      | type  | ticker | quantity |
      | stock | TSLA   | 1.1      |
      | stock | AAPL   | 10       |
    Then the total value should not be 0                                  # wallet_test.go:86 -> theTotalValueShouldNotBeUSD

  Scenario: Get wallet value in EUR with exchange rate from exchangerate.host # features/get_wallet_value.feature:39
    Given I have these assets:                                                # wallet_test.go:28 -> iHaveTheseAssets
      | type   | ticker | quantity |
      | stock  | TSLA   | 1.1      |
      | stock  | AAPL   | 10       |
      | crypto | BTC    | 5        |
    Then the total value should not be 0 EUR                                  # wallet_test.go:97 -> theTotalValueShouldNotBeCurrency

7 scenarios (7 passed)
23 steps (23 passed)
591.004488ms
```

You can also see the output in the [Actions workflow](https://github.com/angristan/go-wallet-bdd-test/actions).
