package main

import "fmt"
import "os"
import "bufio"
import "regexp"

//import "reflect"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	//regex to match input
	re_transac := regexp.MustCompile(`^(?:\p{Sc}\s)?(?:-|\+)?(?:0(?:\.0*|,0*))?([1-9]).*$`)
	m := map[string]int{"1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0}

	for i := 0; i < N; i++ {
		scanner.Scan()
		transaction := scanner.Text()
		s := re_transac.FindStringSubmatch(transaction)

		m[s[len(s)-1]] += 1
	}

	for k, v := range m {
		var p float32 = float32(v) / float32(N)
		var margin float32 = 0.1

		switch k {
		case "1":
			if p < 0.301-margin || p > 0.301+margin {
				fmt.Println("true") // Write answer to stdout
				os.Exit(0)
			}
		case "2":
			if p < 0.176-margin || p > 0.176+margin {
				fmt.Println("true") // Write answer to stdout
				os.Exit(0)
			}
		case "3":
			if p < 0.125-margin || p > 0.125+margin {
				fmt.Println("true") // Write answer to stdout
				os.Exit(0)
			}
		case "4":
			if p < 0.097-margin || p > 0.097+margin {
				fmt.Println("true") // Write answer to stdout
				os.Exit(0)
			}
		case "5":
			if p < 0.079-margin || p > 0.079+margin {
				fmt.Println("true") // Write answer to stdout
				os.Exit(0)
			}
		case "6":
			if p < 0.067-margin || p > 0.067+margin {
				fmt.Println("true") // Write answer to stdout
				os.Exit(0)
			}
		case "7":
			if p < 0.058-margin || p > 0.058+margin {
				fmt.Println("true") // Write answer to stdout
				os.Exit(0)
			}
		case "8":
			if p < 0.051-margin || p > 0.051+margin {
				fmt.Println("true") // Write answer to stdout
				os.Exit(0)
			}
		case "9":
			if p < 0.046-margin || p > 0.046+margin {
				fmt.Println("true") // Write answer to stdout
				os.Exit(0)
			}
		}
	}
	fmt.Println("false") // Write answer to stdout
}
