package pascal

func safeAccess(triangle [][]int, row, col int) int {
	if row < 0 || row >= len(triangle) {
		return 0
	}
	if col < 0 || col >= len(triangle[row]) {
		return 0
	}
	return triangle[row][col]
}

func calcValue(triangle [][]int, row, col int) int {
	return safeAccess(triangle, row-1, col-1) + safeAccess(triangle, row-1, col)
}

func Triangle(rows int) [][]int {
	var triangle [][]int
	triangle = append(triangle, []int{1})
	for i := 1; i < rows; i++ {
		var row []int
		for j := 0; j <= i; j++ {
			row = append(row, calcValue(triangle, i, j))
		}
		triangle = append(triangle, row)
	}
	return triangle
}
