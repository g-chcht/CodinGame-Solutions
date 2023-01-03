package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	mayanBase = 20
)

// using global variables. not pretty but efficient
var H, L int
var numerals = make([]string, mayanBase)

func find(s []string, e string) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

func printMayanNumber(s []string) {
	for _, v := range s {
		for j := 0; j < H; j++ {
			start, end := j*L, j*L+H
			fmt.Println(v[start:end])
		}
	}
}

func mayanToArabic(n []string) int64 {
	indexes := make([]int, len(n))
	for i, v := range n {
		indexes[i] = find(numerals, v)
	}

	var b20 []byte
	for _, v := range indexes {
		b20 = strconv.AppendInt(b20, int64(v), mayanBase)
	}

	i, _ := strconv.ParseInt(string(b20), mayanBase, 64)

	return i
}

func arabicToMayan(i int64) []string {
	b20 := strconv.FormatInt(i, mayanBase)
	res := make([]string, len(b20))
	for j, v := range b20 {
		idx, _ := strconv.ParseInt(string(v), mayanBase, 64)
		res[j] = numerals[idx]
	}
	return res
}

func main() {
	fmt.Scan(&L, &H)

	for i := 0; i < H; i++ {
		var numeral string
		fmt.Scan(&numeral)
		for j := 0; j < 20; j++ {
			start, end := L*j, L*(j+1)
			numerals[j] += numeral[start:end]
		}
	}

	var S1 int
	fmt.Scan(&S1)
	n1 := make([]string, S1/H)
	for i := 0; i < S1; i++ {
		var num1Line string
		fmt.Scan(&num1Line)
		n1[int(i/H)] += num1Line
	}

	var S2 int
	fmt.Scan(&S2)
	n2 := make([]string, S2/H)
	for i := 0; i < S2; i++ {
		var num2Line string
		fmt.Scan(&num2Line)
		n2[int(i/H)] += num2Line
	}

	var operation string
	fmt.Scan(&operation)

	a1 := mayanToArabic(n1)
	a2 := mayanToArabic(n2)
	var r int64

	switch operation {
	case "+":
		r = a1 + a2
	case "-":
		r = a1 - a2
	case "*":
		r = a1 * a2
	case "/":
		r = a1 / a2
	default:
		fmt.Fprintln(os.Stderr, "Ooops")
	}

	rMaya := arabicToMayan(r)

	printMayanNumber(rMaya)
}