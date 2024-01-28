package programingskill

import (
	"strconv"
)

func S1275(moves [][]int) string {
	board := [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}

	for i, v := range moves {
		if i < 4 {
			y := strconv.Itoa(i % 2)
			board[v[0]][v[1]] = y
		} else {
			y := strconv.Itoa(i % 2)
			board[v[0]][v[1]] = y
			for j := 0; j < 3; j++ {
				if board[j][1] == board[j][2] && board[j][1] == board[j][0] && board[j][0] != "-" {
					return summery(y)
				}
				if board[1][j] == board[2][j] && board[1][j] == board[0][j] && board[0][j] != "-" {
					return summery(y)
				}
			}
			if board[1][1] == board[2][2] && board[1][1] == board[0][0] && board[0][0] != "-" {
				return summery(y)
			}
			if board[2][0] == board[1][1] && board[1][1] == board[0][2] && board[1][1] != "-" {
				return summery(y)
			}
		}
	}
	if len(moves)<9{
		return "Pending"
	} else{
		return "Draw"
	}

}

func summery(s string) string {
	if s == "1" {
		return "B"
	}
	return "A"
}
