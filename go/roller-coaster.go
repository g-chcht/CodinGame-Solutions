package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var L, C, N int
	fmt.Scan(&L, &C, &N)

	var cycleQueue = make([]int, N)
	var myhash = make(map[int][2]int)
	var head, remember_head, sum, current_cap, people_count int

	for i := 0; i < N; i++ {
		var pi int
		fmt.Scan(&pi)
		cycleQueue[i] = pi
		people_count += pi
	}

	if people_count <= L {
		fmt.Println(people_count * C) // Write answer to stdout
		os.Exit(0)
	}

	for i := 0; i < C; i++ {
		t, present := myhash[head]
		if present {
			head = t[0]
			sum += t[1]
			continue
		}
		remember_head = head
		for current_cap+cycleQueue[head] <= L {
			current_cap = current_cap + cycleQueue[head]
			head = (head + 1) % N
		}
		myhash[remember_head] = [2]int{head, current_cap}
		sum += current_cap
		current_cap = 0
	}

	//fmt.Fprintln(os.Stderr, myhash)
	//fmt.Fprintln(os.Stderr, len(myhash))
	fmt.Println(sum) // Write answer to stdout
}
