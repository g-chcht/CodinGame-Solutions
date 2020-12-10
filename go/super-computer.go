package main

import "fmt"
import "os"
import "errors"
import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var N int
	fmt.Scan(&N)

	intervals := make([][2]int, N)

	for i := 0; i < N; i++ {
		var J, D int
		fmt.Scan(&J, &D)
		intervals[i] = [2]int{J, J + D}
	}
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })

	answer := 0

	for len(intervals) > 0 {
		i := 0
		min := intervals[0][1]
		for idx, a := range intervals {
			if a[0] < min {
				i = idx
				intervals[idx] = [2]int{0, 0}
			} else {
				break
			}
		}
		intervals = intervals[i+1:]
		answer += 1
	}
	intervals = nil
	fmt.Fprintln(os.Stderr, "Debug msg")
	fmt.Println(answer) // Write answer to stdout
}

func findMin(intervals [][2]int) (min int, e error) {
	if len(intervals) == 0 {
		return 0, errors.New("Cannot detect a minimum value in an empty map")
	}
	// J + D < min
	min = 1001000

	for _, a := range intervals {
		if a[1] < min {
			min = a[1]
		}
	}
	return min, nil
}
