package grains

import "fmt"

const (
	minSquare = 1
	maxSquare = 64
)

var errorOutOfRange error = fmt.Errorf("square number must be between %d-%d (inclusive)", minSquare, maxSquare)

func Square(number int) (uint64, error) {
	if number < minSquare || number > maxSquare {
		return 0, errorOutOfRange
	}
	grains := uint64(1)
	for i := 2; i <= number; i++ {
		grains *= 2
	}
	return grains, nil
}

func Total() uint64 {
	grains := uint64(0)
	for i := 1; i <= maxSquare; i++ {
		squareGrains, err := Square(i)
		if err != nil {
			panic(err)
		}
		grains += squareGrains
	}
	return grains
}
