package sieve

// Source: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#Pseudocode

func Sieve(limit int) []int {
	marked := make([]bool, limit+1) // Initially set to false
	for i := 2; i*i <= limit; i++ {
		if !marked[i] {
			for j := i * i; j <= limit; j += i {
				marked[j] = true
			}
		}
	}
	var out []int
	for i := 2; i <= limit; i++ {
		if !marked[i] {
			out = append(out, i)
		}
	}
	return out
}
