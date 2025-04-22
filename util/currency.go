package util

const (
	RUB = "RUB"
	USD = "USD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case RUB, USD, EUR:
		return true
	}
	return false
}
