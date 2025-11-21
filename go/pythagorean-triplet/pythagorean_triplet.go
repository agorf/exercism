package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var triplets []Triplet
	a, b, c := min, min+1, min+2
	for {
		if a*a+b*b == c*c {
			triplets = append(triplets, Triplet{a, b, c})
		}
		if c < max {
			c++
		} else if b+1 < max {
			b++
			c = b + 1
		} else if a+2 < max {
			a++
			b = a + 1
			c = b + 1
		} else {
			break
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var triplets []Triplet
	for _, triplet := range Range(0, p) {
		if triplet[0]+triplet[1]+triplet[2] == p {
			triplets = append(triplets, triplet)
		}
	}
	return triplets
}
