package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var L, C int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&L, &C)

    S := C*L

    var directions = [4]int{C,1,-C,-1}
    var direction,currentPos, k  int
    cityMap := make([]rune, S)
    var path []string
	var breakerMode, blocked, inverted bool
    var teleporters []int
    
    for i := 0; i < L; i++ {
        scanner.Scan()
        row := scanner.Text()
        for j, c := range row {
            cityMap[i*C+j]=c
            if c == '@' {
                currentPos=i*C+j
            } else if c == 'T' {
                teleporters = append(teleporters, i*C+j)
            }
        }
    }

    for cityMap[currentPos] != '$' {
        switch cityMap[currentPos] {
        case 'S':
            direction = 0
        case 'E':
            direction = 1
        case 'N':
            direction = 2
        case 'W':
            direction = 3
		case 'B':
            breakerMode = !breakerMode
		case 'I':
            inverted = !inverted
        case 'T':
            if teleporters[0] == currentPos {
                currentPos = teleporters[1]
            } else {
                currentPos = teleporters[0]
            }
        }

		nextPos := currentPos + directions[direction]
        switch cityMap[nextPos] {
		case 'X':
			if breakerMode {
				cityMap[nextPos] = ' '
			} else {
				blocked = true
			}
        case '#':
			blocked = true
        }

		if blocked {
            if !inverted {
                for i:=0; i<4;i++{
                    if cityMap[currentPos+directions[i]] != '#' && cityMap[currentPos+directions[i]] != 'X' {
                        direction = i
                        break
                    }
                }
            } else {
                for i:=3; i>=0;i--{
                    if cityMap[currentPos+directions[i]] != '#' && cityMap[currentPos+directions[i]] != 'X' {
                        direction = i
                        break
                    }
                }
            }
			blocked = false
			nextPos = currentPos + directions[direction]
		}

        switch direction {
        case 0:
            path = append(path, "SOUTH")
        case 1:
            path = append(path, "EAST")
        case 2:
            path = append(path, "NORTH")
        case 3:
            path = append(path, "WEST")
        }

        currentPos = nextPos
        k++
        if k >= 200 {
            fmt.Println("LOOP")
            os.Exit(0)
        }
    }
    
    for _,v := range path {
        fmt.Println(v)
    }
}