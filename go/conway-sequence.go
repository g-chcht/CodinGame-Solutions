package main

import "fmt"
import "os"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var R int
	fmt.Scan(&R)

	var L int
	fmt.Scan(&L)

	myslice := []string{strconv.Itoa(R)}

	for i := 1; i < L; i++ {
		var tmp string = myslice[0]
		var count int
		var slice_tmp = make([]string, 0)
		for _, val := range myslice {
			if tmp == val {
				count += 1
			} else {
				slice_tmp = append(slice_tmp, strconv.Itoa(count), tmp)
				count = 1
				tmp = val
			}
		}
		slice_tmp = append(slice_tmp, strconv.Itoa(count), tmp)
		myslice = slice_tmp
	}
	fmt.Fprintln(os.Stderr, "Debug msg")
	fmt.Println(strings.Join(myslice, " "))
}
