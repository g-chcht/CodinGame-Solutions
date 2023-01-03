package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func getScore(w string) int {
	var s int
	for _, v := range w {
		switch v {
		case 'e', 'a', 'i', 'o', 'n', 'r', 't', 'l', 's', 'u':
			s += 1
		case 'd', 'g':
			s += 2
		case 'b', 'c', 'm', 'p':
			s += 3
		case 'f', 'h', 'v', 'w', 'y':
			s += 4
		case 'k':
			s += 5
		case 'j', 'x':
			s += 8
		case 'q', 'z':
			s += 10
		default:
			s += 0
		}

	}
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	var dictionary = make([]string, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		W := scanner.Text()
		dictionary[i] = W
	}
	scanner.Scan()
	LETTERS := scanner.Text()
	_ = LETTERS // to avoid unused error

	var maxScore, idxMaxScore int

	for i, v := range dictionary {
		lettersAvailable := LETTERS
		fit := true
		for _, c := range v {
			idx := strings.IndexRune(lettersAvailable, c)
			if idx == -1 {
				fit = false
				break
			} else {
				s := strings.Split(lettersAvailable, "")
				s[idx] = "-"
				lettersAvailable = strings.Join(s, "")
			}
		}
		if fit {
			score := getScore(v)
			if score > maxScore {
				maxScore = score
				idxMaxScore = i
			}
		}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(dictionary[idxMaxScore]) // Write answer to stdout
}
