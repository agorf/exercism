package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func IsTriangle(a, b, c float64) bool {
	return a > 0 && b > 0 && c > 0 && a+b >= c && b+c >= a && a+c >= b
}

func KindFromSides(a, b, c float64) Kind {
	if !IsTriangle(a, b, c) {
		return NaT
	}
	var k Kind
	switch {
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	case a != b && b != c:
		k = Sca
	}
	return k
}
