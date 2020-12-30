package main

import "fmt"
import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var N int
	fmt.Scan(&N)

	var C int
	fmt.Scan(&C)

	sum := 0

	budgets := make([]int, N)

	for i := 0; i < N; i++ {
		var B int
		fmt.Scan(&B)
		sum += B
		budgets[i] = B
	}

	sort.Ints(budgets)

	if sum < C {
		fmt.Println("IMPOSSIBLE")
	} else {
		for i, v := range budgets {
			average := C / (len(budgets) - i)
			if average > v {
				C -= v
				fmt.Println(v)
			} else {
				C -= average
				fmt.Println(average)
			}
		}
	}
}
