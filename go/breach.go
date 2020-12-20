package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "regexp"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	periodic_el := map[int]string{1: "H", 2: "He", 3: "Li", 4: "Be", 5: "B", 6: "C", 7: "N", 8: "O", 9: "F", 10: "Ne", 11: "Na", 12: "Mg", 13: "Al", 14: "Si", 15: "P", 16: "S", 17: "Cl", 18: "Ar", 19: "K", 20: "Ca", 21: "Sc", 22: "Ti", 23: "V", 24: "Cr", 25: "Mn", 26: "Fe", 27: "Co", 28: "Ni", 29: "Cu", 30: "Zn", 31: "Ga", 32: "Ge", 33: "As", 34: "Se", 35: "Br", 36: "Kr", 37: "Rb", 38: "Sr", 39: "Y", 40: "Zr", 41: "Nb", 42: "Mo", 43: "Tc", 44: "Ru", 45: "Rh", 46: "Pd", 47: "Ag", 48: "Cd", 49: "In", 50: "Sn", 51: "Sb", 52: "Te", 53: "I", 54: "Xe", 55: "Cs", 56: "Ba", 57: "La", 58: "Ce", 59: "Pr", 60: "Nd", 61: "Pm", 62: "Sm", 63: "Eu", 64: "Gd", 65: "Tb", 66: "Dy", 67: "Ho", 68: "Er", 69: "Tm", 70: "Yb", 71: "Lu", 72: "Hf", 73: "Ta", 74: "W", 75: "Re", 76: "Os", 77: "Ir", 78: "Pt", 79: "Au", 80: "Hg", 81: "Tl", 82: "Pb", 83: "Bi", 84: "Po", 85: "At", 86: "Rn", 87: "Fr", 88: "Ra", 89: "Ac", 90: "Th", 91: "Pa", 92: "U", 93: "Np", 94: "Pu", 95: "Am", 96: "Cm", 97: "Bk", 98: "Cf", 99: "Es", 100: "Fm", 101: "Md", 102: "No", 103: "Lr", 104: "Rf", 105: "Db", 106: "Sg", 107: "Bh", 108: "Hs", 109: "Mt", 110: "Ds", 111: "Rg", 112: "Cn", 113: "Nh", 114: "Fl", 115: "Mc", 116: "Lv", 117: "Ts", 118: "Og"}

	for {
		var lines int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &lines)
		raw_lock := make([]string, lines)

		for i := 0; i < lines; i++ {
			scanner.Scan()
			line := scanner.Text()
			raw_lock[i] = line
			fmt.Fprintln(os.Stderr, "line:", line)
		}

		lockType := (strings.Split(raw_lock[0], ":"))[0]
		fmt.Fprintln(os.Stderr, "lockType:", lockType)

		switch lockType {
		case "ss_n":
			fmt.Println(solve_xs_n(lockType, raw_lock))
		case "rs_n":
			fmt.Println(solve_xs_n(lockType, raw_lock))
		case "ss_f":
			re := regexp.MustCompile(`[a-z]`)
			i := re.FindStringIndex(raw_lock[1])[0]
			fmt.Println(raw_lock[1][i] - 'a')
		case "rs_f":
			fmt.Println(raw_lock[1][0] - 'a')
		case "gs_m":
			re := regexp.MustCompile(`\d+`)
			matches := re.FindAllString(raw_lock[1], -1)
			s, _ := strconv.Atoi(matches[len(matches)-1])
			fmt.Println(periodic_el[s])
		case "ss_m":
			re := regexp.MustCompile(`\w+`)
			matches := re.FindAllString(raw_lock[1], -1)
			s := matches[len(matches)-1]
			for k, v := range periodic_el {
				if v == s {
					fmt.Println(k)
					break
				}
			}
		case "ss_asc":
			fmt.Println(solve_ss_asc(raw_lock))
		case "ss_con":
			re := regexp.MustCompile(`¬(\w)\.`)
			matches := re.FindAllStringSubmatch(raw_lock[1], -1)
			fmt.Println(solve_ss_con(matches))
		case "ss_colv":
			re := regexp.MustCompile(`¬(\w)\+`)
			matches := re.FindAllStringSubmatch(raw_lock[1], -1)
			fmt.Fprintln(os.Stderr, matches)
			fmt.Println(solve_xx_colv(matches[0][1]))
		case "rs_colv":
			re := regexp.MustCompile(`¬(\w)`)
			matches := re.FindAllStringSubmatch(raw_lock[1], -1)
			fmt.Fprintln(os.Stderr, matches)
			fmt.Println(solve_xx_colv(matches[0][1]))
		}
	}
}

func solve_xx_colv(letter string) (color string) {
	fmt.Fprintln(os.Stderr, "letter:", letter)
	switch letter {
	case "W":
		color = "GRAY"
	case "w":
		color = "WHITE"
	case "R":
		color = "RED"
	case "r":
		color = "LIGHT_RED"
	case "G":
		color = "GREEN"
	case "g":
		color = "LIGHT_GREEN"
	case "B":
		color = "BLUE"
	case "b":
		color = "LIGHT_BLUE"
	case "y":
		color = "YELLOW"
	case "o":
		color = "ORANGE"
	case "P":
		color = "PINK"
	case "p":
		color = "LIGHT_PINK"
	case "V":
		color = "VIOLET"
	case "v":
		color = "LIGHT_VIOLET"
	case "?":
		color = "CORRUPT"
	default:
		color = "DARK"
	}
	return color
}

func solve_ss_con(m [][]string) (r int) {
	for i, v := range m {
		if v[1] == "r" {
			r = i + 1
			break
		}
	}
	return r
}

func solve_ss_asc(raw_lock []string) (res string) {
	max := -1
	for k := 1; k < len(raw_lock); k++ {
		if len(raw_lock[k]) > max {
			max = len(raw_lock[k])
		}
	}
	for k := 1; k < len(raw_lock); k++ {
		for len(raw_lock[k]) < max {
			raw_lock[k] += " "
		}
	}
	//fmt.Fprintln(os.Stderr, "max:", max)
	myslice := []string{}
	for i := 0; i < max; i++ {
		str := ""
		for j := 1; j < len(raw_lock); j++ {
			c := string(raw_lock[j][i])
			//fmt.Fprintln(os.Stderr, "i:", i, "j:", j, "c:", c)
			str += c
		}
		//fmt.Fprintln(os.Stderr, "str:", "|"+str+"|")
		if strings.Trim(str, " ") == "" {
			//fmt.Fprintln(os.Stderr, "STR IS EMPTY")
			if myslice != nil {
				res += solve_ss_asc_sub(myslice)
				//fmt.Fprintln(os.Stderr, "res:", res)
				myslice = nil
			}
		} else {
			myslice = append(myslice, str)
		}
	}
	return res
}

func solve_ss_asc_sub(number []string) (r string) {
	//fmt.Fprintln(os.Stderr, "number:", number)
	switch number[0] {
	case " +   +":
		switch number[1] {
		case "+++  +":
			r = "2"
		}
	case "+++  +":
		switch number[1] {
		case "+ +  +":
			r = "5"
		}
	case "+     ":
		switch number[1] {
		case "+   ++":
			r = "7"
		}
	case " + +  ":
		switch number[1] {
		case "+ + + ":
			r = "8"
		case "++ ++ ":
			r = "3"
		}
	case " ++++ ":
		switch number[1] {
		case "+ + + ":
			r = "6"
		case "+    +":
			r = "0"
		}
	case " +    ":
		switch number[1] {
		case "+ +   ":
			r = "9"
		case "++   +":
			r = "1"
		}
	case "   +  ":
		switch number[1] {
		case " +++  ":
			r = "4"
		}
	default:
		r = "99599"
	}
	return r
}

func solve_xs_n(lockType string, raw_lock []string) (r int) {

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(raw_lock[1], -1)
	fmt.Fprintln(os.Stderr, "matches:", matches)

	a, _ := strconv.Atoi(matches[0])
	b, _ := strconv.Atoi(matches[1])
	t, _ := strconv.Atoi(matches[len(matches)-1])

	switch lockType {
	case "ss_n":
		r = solve_ss_n(a, b, 1, t)
	case "rs_n":
		r = solve_rs_n(a, b-a, 0, t)
	}
	return r
}

func solve_ss_n(a, b, n, t int) int {
	//fmt.Fprintln(os.Stderr, "a:", a, "b:", b)
	if n == t {
		return b
	}
	return solve_ss_n(b, a+b, n+1, t)
}

func solve_rs_n(a, d, n, t int) int {
	//fmt.Fprintln(os.Stderr, "a:", a, "d:", d, "n:", n, "t:", t)
	if n == t {
		return a
	}
	return solve_rs_n(a+d, d, n+1, t)
}
