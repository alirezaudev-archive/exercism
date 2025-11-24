package stateoftictactoe

import (
	"errors"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	xCount := 0
	oCount := 0

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r][c] == 'X' {
				xCount++
			} else if board[r][c] == 'O' {
				oCount++
			}
		}
	}

	if oCount > xCount || xCount > oCount+1 {
		return "", errors.New("wrong turn order")
	}

	lines := []string{
		board[0], board[1], board[2],
		string([]byte{board[0][0], board[1][0], board[2][0]}),
		string([]byte{board[0][1], board[1][1], board[2][1]}),
		string([]byte{board[0][2], board[1][2], board[2][2]}),
		string([]byte{board[0][0], board[1][1], board[2][2]}),
		string([]byte{board[0][2], board[1][1], board[2][0]}),
	}

	xWins := false
	oWins := false
	for _, line := range lines {
		if line == "XXX" {
			xWins = true
		}
		if line == "OOO" {
			oWins = true
		}
	}

	if (xWins && oWins) ||
		(xWins && xCount == oCount) ||
		oWins && xCount > oCount {
		return "", errors.New("error")
	}

	if xWins || oWins {
		return Win, nil
	}

	if xCount+oCount == 9 {
		return Draw, nil
	}

	return Ongoing, nil
}
