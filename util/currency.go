package util

// Supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	JPY = "JPY"
	GBP = "GBP"
	CNY = "CNY"
	AUD = "AUD"
	CAD = "CAD"
	CHF = "CHF"
	HKD = "HKD"
	RUB = "RUB"
)

// IsSupported returns true if currency is supported
func IsSupported(currency string) bool {
	switch currency {
	case USD, EUR, JPY, GBP, CNY, AUD, CAD, CHF, HKD, RUB:
		return true
	default:
		return false
	}
}
