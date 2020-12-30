package main

import "fmt"
import "os"
import "bufio"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func next(s string) (res string) {
	carry := true
	for i := len(s) - 1; i >= 0; i-- {
		if carry {
			if s[i] == '9' {
				res = "0" + res
				carry = true
			} else {
				res = string(s[i]+1) + res
				carry = false
			}
		} else {
			res = s[:i+1] + res
			break
		}
		//fmt.Fprintln(os.Stderr, "res:", res)
	}

	if carry {
		res = "1" + res
	}

	return res
}

func is_growing(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	n := scanner.Text()
	fmt.Fprintln(os.Stderr, "n:", n)

	s := string(n[0])
	changed := false
	for i := 1; i < len(n); i++ {
		fmt.Fprintln(os.Stderr, "n[i]:", n[i]-'0', "n[i-1]:", n[i-1]-'0')
		if n[i] < s[i-1] {
			s += string(s[i-1])
			changed = true
		} else if n[i] > s[i-1] && changed {
			s += string(s[i-1])
		} else {
			s += string(n[i])
		}
	}

	if !changed {
		n = next(n)

		for !is_growing(n) {
			n = next(n)
			fmt.Fprintln(os.Stderr, "n:", n)
		}
		s = n
	}

	fmt.Println(s) // Write answer to stdout
}
