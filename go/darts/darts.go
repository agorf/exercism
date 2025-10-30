package darts

func Score(x, y float64) int {
	dsq := x*x + y*y
	switch {
	case dsq <= 1*1:
		return 10
	case dsq <= 5*5:
		return 5
	case dsq <= 10*10:
		return 1
	default:
		return 0
	}
}
