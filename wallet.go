package main

var prices map[string]float64
var rates map[string]float64

type asset struct {
	category AssetCategory
	ticker   string
	amount   float64
}

type AssetCategory int

const (
	Stock          AssetCategory = iota
	CryptoCurrency               = iota
	FIAT                         = iota
)

var assets []asset

func main() {
	// TODO, out of scope.
}

func addAsset(a asset) {
	assets = append(assets, a)
}

func getWalletValue(currency string) (float64, error) {
	USDValue := 0.0
	for _, a := range assets {
		assetPrice, err := getAssetPrice(a)
		if err != nil {
			return 0, err
		}
		USDValue += assetPrice * a.amount
	}
	if currency != "USD" {
		return convertUSDto(USDValue, currency)
	}
	return USDValue, nil
}
