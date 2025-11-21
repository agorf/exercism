package minesweeper

func isMine(board []string, row, col int) bool {
	if row < 0 || row >= len(board) {
		return false
	}
	if col < 0 || col >= len(board[0]) {
		return false
	}
	return board[row][col] == '*'
}

func mineCount(board []string, row, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if isMine(board, row+i, col+j) {
				count++
			}
		}
	}
	return count
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	var newBoard []string
	for i := range len(board) {
		var newRow []byte
		for j := range len(board[i]) {
			switch board[i][j] {
			case ' ':
				count := mineCount(board, i, j)
				if count > 0 {
					newRow = append(newRow, byte('0'+count))
					continue
				}
				fallthrough
			default:
				newRow = append(newRow, board[i][j])
			}
		}
		newBoard = append(newBoard, string(newRow))
	}
	return newBoard
}
