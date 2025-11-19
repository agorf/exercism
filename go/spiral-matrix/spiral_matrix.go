package spiralmatrix

const (
	colInc = iota
	rowInc
	colDec
	rowDec
)

func SpiralMatrix(size int) [][]int {
	matrix := make([][]int, 0, size)
	if size == 0 {
		return matrix
	}
	for range size {
		row := make([]int, size)
		matrix = append(matrix, row)
	}
	i, j := 0, 0
	num := 1
	state := colInc
	for {
		matrix[i][j] = num
		num++
		if num > size*size {
			break
		}
		switch state {
		case colInc:
			if j < size-1 && matrix[i][j+1] == 0 {
				j++
			} else {
				state = rowInc
				i++
			}
		case rowInc:
			if i < size-1 && matrix[i+1][j] == 0 {
				i++
			} else {
				state = colDec
				j--
			}
		case colDec:
			if j > 0 && matrix[i][j-1] == 0 {
				j--
			} else {
				state = rowDec
				i--
			}
		case rowDec:
			if i > 0 && matrix[i-1][j] == 0 {
				i--
			} else {
				state = colInc
				j++
			}
		}
	}
	return matrix
}
