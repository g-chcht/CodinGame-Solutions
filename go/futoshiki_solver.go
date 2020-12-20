package main

import "fmt"
import "os"
import "bufio"
import "unicode"
import "regexp"
import "math"
import "io"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type position struct {
	id    int
	value int
	given bool
	inf   []int
}

type board struct {
	positions  []*position
	board_size int
	state      []int
}

func newBoard() *board {
	b := board{make([]*position, 0), 0, []int{0, 0}}
	return &b
}

func newPosition(id int) *position {
	p := position{id, 0, false, []int{}}
	return &p
}

func (p *position) setPosition(value rune) {
	p.value = int(value - '0')
	if p.value != 0 {
		p.given = true
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var size int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &size)

	first := true
	n, row, counter := 0, 0, 0
	myboard := newBoard()

	for i := 0; i < size; i++ {
		scanner.Scan()
		line := scanner.Text()
		//fmt.Fprintln(os.Stderr, "line:", line)
		re := regexp.MustCompile(`\d+`)
		n = len(re.FindAllString(line, -1))

		if first && n > 0 {
			first = false
			myboard.board_size = n
			myboard.positions = make([]*position, n*n)
			for i := 0; i < n*n; i++ {
				myboard.positions[i] = newPosition(i)
			}
		}

		for k, v := range line {
			if unicode.IsDigit(v) {
				myboard.positions[counter].setPosition(v)
				counter += 1
			} else if unicode.IsSymbol(v) && string(v) == ">" {
				left := row*myboard.board_size + offset(k-1)
				right := row*myboard.board_size + offset(k+1)
				myboard.positions[left].inf = append(myboard.positions[left].inf, right)
			} else if unicode.IsSymbol(v) && string(v) == "<" {
				left := row*myboard.board_size + offset(k+1)
				right := row*myboard.board_size + offset(k-1)
				myboard.positions[left].inf = append(myboard.positions[left].inf, right)
			} else if unicode.IsSymbol(v) && string(v) == "^" {
				left := row*myboard.board_size + offset(k)
				right := (row-1)*myboard.board_size + offset(k)
				myboard.positions[left].inf = append(myboard.positions[left].inf, right)
			} else if unicode.IsLetter(v) && string(v) == "v" {
				left := (row-1)*myboard.board_size + offset(k)
				right := row*myboard.board_size + offset(k)
				myboard.positions[left].inf = append(myboard.positions[left].inf, right)
			}
		}
		if n > 0 {
			row += 1
		}
	}

	bt(myboard)
}

func bt(myboard *board) {
	if reject(myboard) {
		return
	}

	if accept(myboard.positions) {
		printBoard(myboard.positions, os.Stdout)
		os.Exit(0)
	}

	first(myboard)

	for {
		bt(myboard)
		next(myboard)
	}
}

func first(myboard *board) {
	for _, p := range myboard.positions {
		if p.value == 0 {
			myboard.state[0] = p.id
			myboard.state[1] = best_val(myboard, p.id)
			myboard.positions[p.id].value = myboard.state[1]
			break
		}
	}
}

func next(myboard *board) {
	myboard.state[1] += 1
	for myboard.state[1] > myboard.board_size {
		myboard.positions[myboard.state[0]].value = 0
		myboard.state[0] -= 1
		for myboard.positions[myboard.state[0]].given {
			myboard.state[0] -= 1
		}

		myboard.state[1] = best_val(myboard, myboard.state[0])
		myboard.positions[myboard.state[0]].value = myboard.state[1]
	}
	myboard.positions[myboard.state[0]].value = myboard.state[1]
}

func best_val(myboard *board, idPos int) int {
	/*
	   row_b_inf := idPos / myboard.board_size
	   row_b_sup := idPos / myboard.board_size + 1
	   id_b_inf := row_b_inf * myboard.board_size
	   id_b_sup := row_b_sup * myboard.board_size
	   m := map[int]bool{}
	   //fmt.Fprintln(os.Stderr, "idPos:", idPos, "row_b_inf:", row_b_inf, "row_b_sup:", row_b_sup, "id_b_inf:", id_b_inf, "id_b_sup:", id_b_sup)
	   for i := id_b_inf; i < id_b_sup; i++ {
	       m[myboard.positions[i].value] = true
	   }
	   temp := idPos % myboard.board_size
	   for temp < myboard.board_size * myboard.board_size {
	       m[myboard.positions[temp].value] = true
	       temp = temp + myboard.board_size
	   }
	   for j := 1; j < myboard.board_size; j++ {
	       _, prs := m[j]
	       if !prs {
	           return j
	       }
	   }
	*/

	return myboard.positions[idPos].value + 1
}

func reject(myboard *board) bool {
	m := make(map[int]int)

	//validate row by row
	for i, v := range myboard.positions {
		if v.value != 0 && found_issue_in_superior(myboard, v) {
			return true
		}
		if v.value != 0 {
			m[v.value] += 1
		}
		if (i+1)%myboard.board_size == 0 {
			for key, val := range m {
				if val > 1 {
					return true
				} else {
					delete(m, key)
				}
			}
		}
	}

	//validate column by column
	for i := 0; i < myboard.board_size; i++ {
		for j := 0; j < myboard.board_size; j++ {
			v := myboard.positions[i+j*myboard.board_size]
			if v.value != 0 {
				m[v.value] += 1
			}
		}
		for key, val := range m {
			if val > 1 {
				return true
			} else {
				delete(m, key)
			}
		}
	}
	return false
}

func found_issue_in_superior(myboard *board, p *position) bool {
	myId := p.id
	myValue := p.value

	for _, otherId := range myboard.positions[myId].inf {
		otherValue := myboard.positions[otherId].value
		if myValue <= otherValue {
			return true
		}
	}
	return false
}

func accept(board []*position) bool {

	size := int(math.Sqrt(float64(len(board))))
	m := make(map[int]int)

	//validate row by row
	for i, v := range board {
		if v.value != 0 {
			m[v.value] += 1
		} else {
			return false
		}
		if (i+1)%size == 0 {
			for key, val := range m {
				if val > 1 {
					return false
				} else {
					delete(m, key)
				}
			}
		}
	}

	//validate column by column
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			v := board[i+j*size]
			if v.value != 0 {
				m[v.value] += 1
			} else {
				return false
			}
		}
		for key, val := range m {
			if val > 1 {
				return false
			} else {
				delete(m, key)
			}
		}
	}
	return true
}

func offset(n int) (r int) {
	switch n {
	case 0:
		r = 0
	default:
		r = n / 2
	}
	return r
}

func printBoard(board []*position, w io.Writer) {
	board_size := int(math.Sqrt(float64(len(board))))
	for i, v := range board {
		fmt.Fprint(w, v.value)
		if w == os.Stderr {
			fmt.Fprint(w, "|")
		}
		if (i+1)%board_size == 0 {
			fmt.Fprint(w, "\n")
		}
	}
}
