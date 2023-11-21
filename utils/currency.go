package utils

const (
	// USD is the currency of the United States dollar
	USD = "USD"
	// EUR is the currency of the Euro
	EUR = "EUR"
	// CAD is the currency of the Canadian dollar
	CAD = "CAD"
)

// IsSupportedCurrency checks if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
