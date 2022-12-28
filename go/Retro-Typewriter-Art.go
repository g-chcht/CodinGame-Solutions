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

	for _, v := range strings.Split(scanner.Text(), " ") {
		printInstr(v)
	}
}

func printInstr(instr string) {
	switch {
	case strings.HasSuffix(instr, "sp"):
		n, _ := strconv.Atoi(instr[:len(instr)-2])
		fmt.Print(strings.Repeat(" ", n))
	case strings.HasSuffix(instr, "bS"):
		n, _ := strconv.Atoi(instr[:len(instr)-2])
		fmt.Print(strings.Repeat("\\", n))
	case strings.HasSuffix(instr, "sQ"):
		n, _ := strconv.Atoi(instr[:len(instr)-2])
		fmt.Print(strings.Repeat("'", n))
	case strings.HasSuffix(instr, "nl"):
		fmt.Print("\n")
	default:
		n, _ := strconv.Atoi(instr[0 : len(instr)-1])
		c := instr[len(instr)-1:]
		fmt.Print(strings.Repeat(c, n))
	}
}
