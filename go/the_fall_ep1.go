package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // W: number of columns.
    // H: number of rows.
    var W, H int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&W, &H)

	rooms := make([][]string, H)
    
    for i := 0; i < H; i++ {
        scanner.Scan()
        LINE := scanner.Text()
		rooms[i] = strings.Fields(LINE)
	}

    var EX int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&EX)
    for {
        var XI, YI int
        var POS string
        scanner.Scan()
        fmt.Sscan(scanner.Text(),&XI, &YI, &POS)

		roomType := rooms[YI][XI]
		fmt.Fprintf(os.Stderr, "XI: %d, YI:%d, roomType: %s\n", XI, YI, roomType)

		switch roomType {
		case "1", "3", "7", "8", "9", "12", "13":
			fmt.Println(XI, YI+1)
		case "11":
			fmt.Println(XI+1, YI)
		case "10":
			fmt.Println(XI-1, YI)
		case "2", "6":
			if POS == "LEFT"{
				fmt.Println(XI+1, YI)
			} else {
				fmt.Println(XI-1, YI)
			}
		case "4":
			if POS == "TOP"{
				fmt.Println(XI-1, YI)
			} else {
				fmt.Println(XI, YI+1)
			}
		case "5":
			if POS == "LEFT"{
				fmt.Println(XI, YI+1)
			} else {
				fmt.Println(XI+1, YI)
			}
		}
		}
}