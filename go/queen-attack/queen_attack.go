package queenattack

import "errors"

const (
	minRow = '1'
	maxRow = '8'
	minCol = 'a'
	maxCol = 'h'
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Queen struct {
	row, col int
}

func (q Queen) IsEqual(other Queen) bool {
	return q.row == other.row && q.col == other.col
}

func (q Queen) CanAttack(other Queen) bool {
	return q.row == other.row ||
		q.col == other.col ||
		abs(q.row-other.row) == abs(q.col-other.col) // Diagonal
}

func isInvalidPosition(position string) bool {
	return position == "" ||
		len(position) != 2 ||
		position[0] < minCol ||
		position[0] > maxCol ||
		position[1] < minRow ||
		position[1] > maxRow
}

func NewQueen(position string) (Queen, error) {
	if isInvalidPosition(position) {
		return Queen{}, errors.New("invalid position")
	}
	q := Queen{
		row: int(position[0] - minCol + 1),
		col: int(position[1] - minRow + 1),
	}
	return q, nil
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	white, err := NewQueen(whitePosition)
	if err != nil {
		return false, err
	}
	black, err := NewQueen(blackPosition)
	if err != nil {
		return false, err
	}
	if white.IsEqual(black) {
		return false, errors.New("positions should differ")
	}
	return white.CanAttack(black), nil
}
