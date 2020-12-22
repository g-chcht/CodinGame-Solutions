package main

import "fmt"
import "os"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H int
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	fmt.Scan(&X0, &Y0)

	x, y, start, end := X0, Y0, 0, W-1
	xIsDone := false
	z, z_old, Z := x, x, W

	for {
		// bombDir: Current distance to the bomb compared to previous distance (COLDER, WARMER, SAME or UNKNOWN)
		var bombDir string
		fmt.Scan(&bombDir)

		fmt.Fprintln(os.Stderr, "bombDir:", bombDir)

		start, end, z, z_old = binary_search(start, end, z, z_old, Z, bombDir)
		if start == end && !xIsDone {
			x = z
			xIsDone = true
			z, z_old, Z = y, y, H
			start, end = 0, H-1
		}

		if !xIsDone {
			x = z
		} else {
			y = z
		}
		fmt.Println(x, y)
	}
}

func binary_search(start, end, z, z_old, Z int, bombDir string) (int, int, int, int) {

	switch bombDir {
	case "UNKNOWN":
		//Do nothing
	case "WARMER":
		d := int(math.Abs(float64(z - z_old)))
		if z > z_old {
			tmp := z_old + d/2 + 1
			if tmp > start {
				start = tmp
			}
		} else if z < z_old {
			tmp := 0
			if d%2 == 0 {
				tmp = z + d/2 - 1
			} else {
				tmp = z + d/2
			}
			if tmp < end {
				end = tmp
			}
			if end < start {
				end = start
			}
		}
	case "COLDER":
		d := int(math.Abs(float64(z - z_old)))
		if z > z_old {
			if d%2 == 0 {
				end = z_old + d/2 - 1
			} else {
				end = z_old + d/2
			}
			if end < start {
				end = start
			}
		} else if z < z_old {
			start = z + d/2 + 1
		}
	case "SAME":
		{
			d := int(math.Abs(float64(z - z_old)))
			if z > z_old {
				start, end = z_old+d/2, z_old+d/2
			} else if z_old > z {
				start, end = z+d/2, z+d/2
			}
			z_old = z
			z = start

			return start, end, z, z_old
		}
	}

	if start == end {
		return start, end, start, z_old
	}

	next_hop := find_next_hop(start, end, z, Z)
	z_old = z
	z = next_hop

	return start, end, z, z_old
}

func find_next_hop(start, end, z, Z int) (next_hop int) {

	mid := find_mid(start, end)
	d := 0
	if (end-start)%2 == 0 {
		d = int(math.Abs(float64(mid - z)))
	} else {
		d = int(math.Abs(float64(mid-z))) - 1
		if d < 0 {
			d = 0
		}
	}

	if z == 0 {
		next_hop = mid
	} else if z > mid {
		next_hop = mid - d
		if next_hop < 0 {
			k := z / 2
			if k < end {
				next_hop = start
			} else {
				next_hop = end
			}
		}
	} else if z <= mid {
		next_hop = mid + d
		if next_hop >= Z {
			next_hop = end
		}
	}

	if next_hop == z {
		next_hop += 1
	}

	return next_hop
}

func find_mid(start, end int) int {
	return start + (end-start)/2
}
