package stateoftictactoe

import "errors"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
	blank   byte  = ' '
)

func countMoves(board []string, player byte) int {
	count := 0
	for i := range 3 {
		for j := range 3 {
			if board[i][j] == player {
				count++
			}
		}
	}
	return count
}

func boardColWin(board []string, col int) (byte, bool) {
	return board[0][col], board[0][col] != blank && board[0][col] == board[1][col] && board[1][col] == board[2][col]
}

func boardRowWin(board []string, row int) (byte, bool) {
	return board[row][0], board[row][0] != blank && board[row][0] == board[row][1] && board[row][1] == board[row][2]
}

func boardBackSlashWin(board []string) (byte, bool) {
	return board[0][0], board[0][0] != blank && board[0][0] == board[1][1] && board[1][1] == board[2][2]
}

func boardSlashWin(board []string) (byte, bool) {
	return board[2][0], board[2][0] != blank && board[2][0] == board[1][1] && board[1][1] == board[0][2]
}

func boardWinners(board []string) map[byte]bool {
	winners := make(map[byte]bool)
	var winner byte
	var won bool
	for x := range 3 {
		winner, won = boardColWin(board, x)
		if won {
			winners[winner] = true
		}
		winner, won = boardRowWin(board, x)
		if won {
			winners[winner] = true
		}
	}
	winner, won = boardBackSlashWin(board)
	if won {
		winners[winner] = true
	}
	winner, won = boardSlashWin(board)
	if won {
		winners[winner] = true
	}
	return winners
}

func isOngoingGame(board []string) bool {
	for i := range 3 {
		for j := range 3 {
			if board[i][j] == blank {
				return true
			}
		}
	}
	return false
}

func StateOfTicTacToe(board []string) (State, error) {
	xMoves := countMoves(board, 'X')
	oMoves := countMoves(board, 'O')
	if oMoves > xMoves {
		return "", errors.New("player O started")
	}
	if xMoves-oMoves > 1 {
		return "", errors.New("player X went twice")
	}
	winners := len(boardWinners(board))
	if winners > 1 {
		return "", errors.New("player X won and O kept playing")
	}
	switch {
	case winners == 1:
		return Win, nil
	case isOngoingGame(board):
		return Ongoing, nil
	default:
		return Draw, nil
	}
}
