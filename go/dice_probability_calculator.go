package main

import "fmt"
import "os"
import "bufio"
import "strings"

//import "math"
import "strconv"
import "regexp"
import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	expr := scanner.Text()
	fmt.Fprintln(os.Stderr, "expr:", expr)

	expr_exploded := explode(expr) // add spaces
	fmt.Fprintln(os.Stderr, "expr after explode:", expr_exploded)

	list_rpn := shunting_yard(expr_exploded)

	m := make(map[int]int)
	for _, v := range *list_rpn {
		m[solve(v)] += 1
	}
	//fmt.Fprintln(os.Stderr, m)

	si := make([]int, 0, len(m))
	for i := range m {
		si = append(si, i)
	}
	sort.Ints(si)
	for _, e := range si {
		fmt.Printf("%d %.2f\n", e, (float64(m[e])/float64(len(*list_rpn)))*100)
	}
}

var opa = map[string]struct {
	prec   int
	rAssoc bool
}{
	"*": {3, false},
	"+": {2, false},
	"-": {2, false},
	">": {1, false},
}

func explode(expr string) string {

	re_op := regexp.MustCompile(`(\+|\-|\>|\*)`)
	expr = re_op.ReplaceAllString(expr, " $1 ")
	//fmt.Fprintln(os.Stderr, "expr after re_op:", expr)

	re_lp := regexp.MustCompile(`\((\d+)`)
	expr = re_lp.ReplaceAllString(expr, "( $1")
	//fmt.Fprintln(os.Stderr, "expr after re_lp:", expr)

	re_rp := regexp.MustCompile(`(\d+)\)`)
	expr = re_rp.ReplaceAllString(expr, "$1 )")
	//fmt.Fprintln(os.Stderr, "expr after re_rp:", expr)

	return expr
}

//Shunting-yard algorithm
func shunting_yard(expr string) *[]string {
	var stack []string // holds operators and left
	list_rpn := []string{""}

	re := regexp.MustCompile(`^d(\d+)$`) // match dice throw

	for _, tok := range strings.Fields(expr) {
		el := string(tok)
		dt := re.FindStringSubmatch(el)

		switch {
		case dt != nil:
			n, _ := strconv.Atoi(dt[1])

			for i, _ := range list_rpn {
				rpn := list_rpn[i]
				list_rpn[i] += " 1" // add operator to result
				for j := 2; j <= n; j++ {
					list_rpn = append(list_rpn, rpn+" "+strconv.Itoa(j))
				}
			}
		case el == "(":
			stack = append(stack, el) // push "(" to stack
		case el == ")":
			var op string
			for {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				//rpn += " " + op // add operator to result
				for i, _ := range list_rpn {
					list_rpn[i] += " " + op // add operator to result
				}
			}
		default:
			if o1, isOp := opa[el]; isOp {
				// token is an operator
				for len(stack) > 0 {
					// consider top item on stack
					op := stack[len(stack)-1]
					if o2, isOp := opa[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}
					// top item is an operator that needs to come off
					stack = stack[:len(stack)-1] // pop it

					for i, _ := range list_rpn {
						list_rpn[i] += " " + op // add operator to result
					}
				}
				// push operator (the new one) to stack
				stack = append(stack, el)
			} else { // element is an operand
				for i, _ := range list_rpn {
					if list_rpn[i] != "" {
						list_rpn[i] += " "
					}
					list_rpn[i] += el // add operand to result
				}
			}
		}
	}
	// drain stack to result
	for len(stack) > 0 {
		//rpn += " " + stack[len(stack)-1]
		for i, _ := range list_rpn {
			list_rpn[i] += " " + stack[len(stack)-1] // add operator to result
		}
		stack = stack[:len(stack)-1]
	}
	return &list_rpn
}

func solve(rpn string) int {
	fmt.Fprintf(os.Stderr, "For postfix %q\n", rpn)
	fmt.Fprintln(os.Stderr, "\nToken            Action            Stack")
	var stack []int
	for _, tok := range strings.Fields(rpn) {
		action := "Apply op to top of stack"
		switch tok {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		/*case "^":
		  stack[len(stack)-2] =
		      math.Pow(stack[len(stack)-2], stack[len(stack)-1])
		  stack = stack[:len(stack)-1]*/
		case ">":
			if stack[len(stack)-2] > stack[len(stack)-1] {
				stack[len(stack)-2] = 1
			} else {
				stack[len(stack)-2] = 0
			}
			stack = stack[:len(stack)-1]
		default:
			action = "Push num onto top of stack"
			//f, _ := strconv.ParseFloat(tok, 64)
			f, _ := strconv.Atoi(tok)
			stack = append(stack, f)
		}
		fmt.Fprintf(os.Stderr, "%3s    %-26s  %v\n", tok, action, stack)
	}
	fmt.Fprintln(os.Stderr, "\nThe final value is", stack[0])
	return stack[0]
}
