package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "regexp"

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

	invalids := make([]string, 0)
	re_ISBN10 := regexp.MustCompile("^([0-9]{9})([0-9]|X)$")
	re_ISBN13 := regexp.MustCompile("^[0-9]{13}$")

	for i := 0; i < N; i++ {
		scanner.Scan()
		ISBN := scanner.Text()
		//fmt.Fprintln(os.Stderr, "ISBN: ", ISBN)

		switch {
		case re_ISBN10.MatchString(ISBN) == true:
			//fmt.Fprintln(os.Stderr, "ISBN 10")
			sum := 0
			weight := 10
			for ind, c := range ISBN {
				//fmt.Fprintln(os.Stderr, "char: ", string(c))
				//skip last
				if ind == len(ISBN)-1 {
					continue
				}
				i, _ := strconv.Atoi(string(c))
				sum += i * weight
				weight -= 1
			}
			//fmt.Fprintln(os.Stderr, "sum: ", sum)
			//fmt.Fprintln(os.Stderr, "sum % 11: ", sum % 11)
			r := strconv.Itoa(11 - sum%11)
			if r == "10" {
				r = "X"
			} else if r == "11" {
				r = "0"
			}
			//fmt.Fprintln(os.Stderr, "r: ", r)
			//fmt.Fprintln(os.Stderr, "ISBN[len(ISBN)-1:]: ", ISBN[len(ISBN)-1:])
			if ISBN[len(ISBN)-1:] != r {
				invalids = append(invalids, ISBN)
				continue
			}
		case re_ISBN13.MatchString(ISBN) == true:
			sum := 0
			weight := 1
			for ind, c := range ISBN {
				//fmt.Fprintln(os.Stderr, "char: ", string(c))
				//skip last
				if ind == len(ISBN)-1 {
					continue
				}
				i, _ := strconv.Atoi(string(c))
				sum += i * weight
				if weight == 1 {
					weight = 3
				} else {
					weight = 1
				}
			}
			/*
			   if ISBN == "9780133661750"{
			       fmt.Fprintln(os.Stderr, "sum: ", sum)
			   }
			*/
			r := strconv.Itoa(10 - sum%10)
			if r == "10" {
				r = "0"
			}
			/*
			   if ISBN == "9780133661750"{
			       fmt.Fprintln(os.Stderr, "r: ", r)
			   }
			*/
			if ISBN[len(ISBN)-1:] != r {
				invalids = append(invalids, ISBN)
				continue
			}
		default:
			//fmt.Fprintln(os.Stderr, "default: ", ISBN)
			invalids = append(invalids, ISBN)
			continue
		}

	}

	fmt.Println(len(invalids), "invalid:") // Write answer to stdout
	for _, v := range invalids {
		fmt.Println(v)
	}
}
