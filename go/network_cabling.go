package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
    var N int
    fmt.Scan(&N)

    var xMin, xMax, sum int
	y := make([]int, N)
    
    for i := 0; i < N; i++ {
        var X, Y int
        fmt.Scan(&X, &Y)
        if i == 0{
            xMin = X
            xMax = X
        } else {
            if xMin > X {
                xMin = X
            }
            if xMax < X {
                xMax = X
            }
        }
		y[i] = Y
    }

    xDiff := xMax - xMin
	sort.Ints(y)
	med := y[N/2]
	for i:=0; i<N;i++ {
		sum += int(math.Abs(float64(y[i]-med)))
	}    
    fmt.Println(xDiff+sum)
}