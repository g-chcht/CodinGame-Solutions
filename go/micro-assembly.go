package main

import "fmt"
import "os"
import "bufio"
import "regexp"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var a, b, c, d int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &a, &b, &c, &d)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	var re = regexp.MustCompile(`^(\w{3})\s(\w+)\s(-?\w+)\s?(-?\w+)?`)
	instructions := make([]string, n)
	var current int

	for i := 0; i < n; i++ {
		scanner.Scan()
		instruction := scanner.Text()
		instructions[i] = instruction
	}

	for current < n {
		var instruction string = instructions[current]
		myslice := re.FindStringSubmatch(instruction)

		switch myslice[1] {
		case "MOV":
			switch myslice[2] {
			case "a":
				a = stringToRes(myslice[3], a, b, c, d)
			case "b":
				b = stringToRes(myslice[3], a, b, c, d)
			case "c":
				c = stringToRes(myslice[3], a, b, c, d)
			case "d":
				d = stringToRes(myslice[3], a, b, c, d)
			}
			current += 1
		case "ADD":
			switch myslice[2] {
			case "a":
				a = stringToRes(myslice[3], a, b, c, d) + stringToRes(myslice[4], a, b, c, d)
			case "b":
				b = stringToRes(myslice[3], a, b, c, d) + stringToRes(myslice[4], a, b, c, d)
			case "c":
				c = stringToRes(myslice[3], a, b, c, d) + stringToRes(myslice[4], a, b, c, d)
			case "d":
				d = stringToRes(myslice[3], a, b, c, d) + stringToRes(myslice[4], a, b, c, d)
			}
			current += 1
		case "SUB":
			switch myslice[2] {
			case "a":
				a = stringToRes(myslice[3], a, b, c, d) - stringToRes(myslice[4], a, b, c, d)
			case "b":
				b = stringToRes(myslice[3], a, b, c, d) - stringToRes(myslice[4], a, b, c, d)
			case "c":
				c = stringToRes(myslice[3], a, b, c, d) - stringToRes(myslice[4], a, b, c, d)
			case "d":
				d = stringToRes(myslice[3], a, b, c, d) - stringToRes(myslice[4], a, b, c, d)
			}
			current += 1
		case "JNE":
			tmp_left := strconv.Itoa(stringToRes(myslice[3], a, b, c, d))
			tmp_right := strconv.Itoa(stringToRes(myslice[4], a, b, c, d))
			if tmp_left != tmp_right {
				current, _ = strconv.Atoi(myslice[2])
			} else {
				current += 1
			}
		}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(a, b, c, d) // Write answer to stdout
}

func stringToRes(str string, a, b, c, d int) int {
	switch str {
	case "a":
		return a
	case "b":
		return b
	case "c":
		return c
	case "d":
		return d
	default:
		myres, _ := strconv.Atoi(str)
		return myres
	}
}
