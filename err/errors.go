package err

import "errors"

var (
	IdCanNotBeMinorThanOneError = errors.New("id can not be minor than zero")

	PriceCanNotBeMinorThanZeroError = errors.New("price can not be minor than zero")

	ErrorTakingCurrencies = errors.New("error taking currencies")
)
