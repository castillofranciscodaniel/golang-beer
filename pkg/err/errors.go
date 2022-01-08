package err

import "errors"

var (
	IdCanNotBeMinorThanOne = errors.New("id can not be minor than zero")

	PriceCanNotBeMinorThanZero = errors.New("price can not be minor than zero")
)
