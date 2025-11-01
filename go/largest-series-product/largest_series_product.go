package lsproduct

import "errors"

func digitsProduct(digits string) (int64, error) {
	var product int64 = 1
	for _, r := range digits {
		if r < '0' || r > '9' {
			return 0, errors.New("digits input must only contain digits")
		}
		product *= int64(r - '0')
		if product == 0 { // Early return
			return 0, nil
		}
	}
	return product, nil
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var maxProduct int64 = 0
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if len(digits) < span {
		return 0, errors.New("span must be smaller than string length")
	}
	for i := 0; i <= len(digits)-span; i++ {
		series := digits[i : i+span]
		product, err := digitsProduct(series)
		if err != nil {
			return 0, err
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct, nil
}
