package perfect

import "errors"

type Classification int

const (
	ClassificationDeficient = iota
	ClassificationPerfect
	ClassificationAbundant
)

type seen struct{}

var (
	seenValue       = seen{}
	ErrOnlyPositive = errors.New("number should be positive")
)

func numberFactors(n int64) []int64 {
	seenFactors := map[int64]seen{}
	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			seenFactors[i] = seenValue
			seenFactors[n/i] = seenValue
		}
	}
	var factors []int64
	for factor := range seenFactors {
		if factor < n { // Exclude self
			factors = append(factors, factor)
		}
	}
	return factors
}

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	var aliquotSum int64
	for _, factor := range numberFactors(n) {
		aliquotSum += factor
	}
	switch {
	case n > aliquotSum:
		return ClassificationDeficient, nil
	case n < aliquotSum:
		return ClassificationAbundant, nil
	default: // n == aliquotSum
		return ClassificationPerfect, nil
	}
}
