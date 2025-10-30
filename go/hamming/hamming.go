package hamming

import "errors"

func Distance(a, b string) (int, error) {
	ar := []rune(a)
	br := []rune(b)

	if len(ar) != len(br) {
		return 0, errors.New("sequences are of different length")
	}

	distance := 0
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}
	return distance, nil
}
