package prime

import "errors"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
// Dummy, slow implementation since another exercise follows that implements the sieve of Eratosthenes
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("'n' is equal or less than zero")
	}
	p := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			p++
			if p == n {
				return i, nil
			}
		}
	}
}
