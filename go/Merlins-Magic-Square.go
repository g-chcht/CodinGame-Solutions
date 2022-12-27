package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	row1 := scanner.Text()
	r := row1
	scanner.Scan()
	row2 := scanner.Text()
	r = r + row2
	scanner.Scan()
	row3 := scanner.Text()
	r = r + row3
	scanner.Scan()
	allButtonsPressed := scanner.Text()

	board := strings.Split(strings.ReplaceAll(r, " ", ""), "")
	solved := []string{"*", "*", "*", "*", "~", "*", "*", "*", "*"}

	for _, v := range allButtonsPressed {
		p, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("Something happened")
			os.Exit(1)
		}
		doNextMove(board, p)
	}

	for i := 1; i <= 9; i++ {
		nextState := make([]string, len(board))
		_ = copy(nextState, board)
		doNextMove(nextState, i)
		if sliceEqual(nextState, solved) {
			fmt.Println(i)
			break
		}
	}
}

func doNextMove(board []string, p int) {
	switch p {
	case 1, 3, 7, 9:
		corner(board, p)
	case 2, 4, 6, 8:
		side(board, p)
	case 5:
		middle(board)
	default:
		fmt.Println("Something happened")
		os.Exit(1)
	}
}

func corner(board []string, p int) {
	switch p {
	case 1:
		board[0] = reverse(board[0])
		board[1] = reverse(board[1])
		board[3] = reverse(board[3])
		board[4] = reverse(board[4])
	case 3:
		board[1] = reverse(board[1])
		board[2] = reverse(board[2])
		board[4] = reverse(board[4])
		board[5] = reverse(board[5])
	case 7:
		board[3] = reverse(board[3])
		board[4] = reverse(board[4])
		board[6] = reverse(board[6])
		board[7] = reverse(board[7])
	case 9:
		board[4] = reverse(board[4])
		board[5] = reverse(board[5])
		board[7] = reverse(board[7])
		board[8] = reverse(board[8])
	}
}

func side(board []string, p int) {
	switch p {
	case 2:
		board[0] = reverse(board[0])
		board[1] = reverse(board[1])
		board[2] = reverse(board[2])
	case 4:
		board[0] = reverse(board[0])
		board[3] = reverse(board[3])
		board[6] = reverse(board[6])
	case 6:
		board[2] = reverse(board[2])
		board[5] = reverse(board[5])
		board[8] = reverse(board[8])
	case 8:
		board[6] = reverse(board[6])
		board[7] = reverse(board[7])
		board[8] = reverse(board[8])
	}
}

func middle(board []string) {
	board[1] = reverse(board[1])
	board[3] = reverse(board[3])
	board[4] = reverse(board[4])
	board[5] = reverse(board[5])
	board[7] = reverse(board[7])
}

func reverse(c string) string {
	if c == "*" {
		return "~"
	} else {
		return "*"
	}
}

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
