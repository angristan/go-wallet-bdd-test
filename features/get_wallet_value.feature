Feature: get wallet value
  Get value of all assets in wallet

  Scenario Outline: Get wallet value in target currency with specified rate
    Given I have these assets:
      | type  | ticker | quantity |
      | stock | TSLA   | 1.1      |
      | stock | AAPL   | 10       |
    And "TSLA" is worth 1000 USD
    And "AAPL" is worth 130 USD
    And "<Currency>" rate is <Rate>
    Then the total value should be <Value> "<Currency>"

    Examples:
      | Currency | Rate   | Value  |
      | EUR      | 0.83   | 1992   |
      | JPY      | 104.95 | 251880 |
      | USD      | 1.0    | 2400   |

  Scenario: Get wallet value in invalid currency
    Given I have these assets:
      | type  | ticker | quantity |
      | stock | TSLA   | 1.1      |
      | stock | AAPL   | 10       |
    Then the total value in "LOL" should show an error

  Scenario: Get wallet value without assets
    Given I have these assets:
      | type | ticker | quantity |
    Then the total value should be 0 USD

  Scenario: Get wallet value in USD with stocks values from Yahoo Finance
    Given I have these assets:
      | type  | ticker | quantity |
      | stock | TSLA   | 1.1      |
      | stock | AAPL   | 10       |
    Then the total value should not be 0

  Scenario: Get wallet value in EUR with exchange rate from exchangerate.host
  and asset values from Yahoo Finance
    Given I have these assets:
      | type   | ticker | quantity |
      | stock  | TSLA   | 1.1      |
      | stock  | AAPL   | 10       |
      | crypto | BTC    | 5        |
    Then the total value should not be 0 EUR
