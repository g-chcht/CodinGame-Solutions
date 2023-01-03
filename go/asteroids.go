package main

import (
	"fmt"
	"os"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type Asteroid struct {
	x1, y1, x2, y2, x3, y3 int
}

func main() {
	var W, H, T1, T2, T3 int
	fmt.Scan(&W, &H, &T1, &T2, &T3)

	m := make(map[byte]*Asteroid)
	var timeCoeff float64 = float64(T3-T2) / float64(T2-T1)
	res := make([]string, 0, H)
	for i := 0; i < H; i++ {
		res = append(res, strings.Repeat(".", W))
	}

	for i := 0; i < H; i++ {
		var firstPictureRow, secondPictureRow string
		fmt.Scan(&firstPictureRow, &secondPictureRow)
		fmt.Fprint(os.Stderr, firstPictureRow, " ")
		fmt.Fprintln(os.Stderr, secondPictureRow)

		for j := 0; j < W; j++ {
			var array = [2]byte{firstPictureRow[j], secondPictureRow[j]}

			for k, v := range array {
				if v != '.' {
					ast, f := m[v]
					if !f {
						if k == 0 {
							m[v] = &Asteroid{x1: j, y1: i}
						} else {
							m[v] = &Asteroid{x2: j, y2: i}
						}
					} else {
						if k == 0 {
							ast.x1 = j
							ast.y1 = i
						} else {
							ast.x2 = j
							ast.y2 = i
						}

						x3Float := float64(ast.x2) + float64(ast.x2-ast.x1)*timeCoeff
						y3Float := float64(ast.y2) + float64(ast.y2-ast.y1)*timeCoeff

                        // out of bounds
                        if x3Float <0 || x3Float >= float64(W) || y3Float <0 || y3Float >= float64(H) {
                            continue
                        }

						ast.x3 = int(x3Float)
						ast.y3 = int(y3Float)
                        

                        c := res[ast.y3][ast.x3]
                        
						if (c != '.') && (c < v) {
							continue
						}
                        res[ast.y3] = replaceAtIndex(res[ast.y3], v, ast.x3)
					}

				}
			}
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

func replaceAtIndex(in string, r byte, i int) string {
	out := []byte(in)
    if i >= 0 && i < len(out){
        out[i] = r
    }
	return string(out)
}
