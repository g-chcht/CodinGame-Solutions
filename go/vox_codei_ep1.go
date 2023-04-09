package main
// code is awful :(

import (
	"fmt"
	"os"
	"strings"
)

var directions = [4]string{"U", "D", "L", "R"}

type Grid struct {
	width, height int
	grid          []string
}

func (g Grid) Copy() *Grid {
	newGrid := make([]string, len(g.grid))
	for i,v := range g.grid {
		newGrid[i] = v
	}
	ng := Grid{width: g.width, height: g.height, grid: newGrid}
	return &ng
}

func (g Grid) Update(pos int, force bool) {
	size := g.height * g.width
	var target int
	min := pos - (pos % g.width)
	max := min + g.width - 1
	t := g.grid[pos]
	
	for i, c := range g.grid {
		switch c {
		case "3":
			g.grid[i] = "2"
		case "2":
			g.grid[i] = "1"
		case "1":
			g.grid[i] = "."
		}
	}

	if t != "." && !force {
		return
	}

	g.grid[pos] = "3"

	for _, d := range directions {
		out:
		for k := 1; k <= 3; k++ {
			switch d {
			case "U":
				target = pos - k*g.width
				if target < 0 {
					break out
				}
			case "D":
				target = pos + k*g.width
				if target >= size {
					break out
				}
			case "L":
				target = pos - k
				if target < min {
					break out
				}
			case "R":
				target = pos + k
				if target > max {
					break out
				}
			}

			if g.grid[target] == "#" {
				break
			} else if g.grid[target] == "@" {
				g.grid[target] = "3"
			}

		}
	}
}

func (g Grid) getNbNodesLeft() int {
	var count int
	for _, c := range g.grid {
		if c == "@" {
			count++
		}
	}
	return count
}

func (g Grid) getNbNodesAlone() int {
	var count int
	size := g.height*g.width
	for i, c := range g.grid {
		if c == "@" {
			alone := true
			target := -1
			pos := i
			min := pos - (pos % g.width)
			max := min + g.width - 1
			for _, d := range directions {
				out:
				for k := 1; k <= 6; k++ {
					switch d {
					case "U":
						target = pos - k*g.width
						if target < 0 {
							break out
						}
					case "D":
						target = pos + k*g.width
						if target >= size {
							break out
						}
					case "L":
						target = pos - k
						if target < min {
							break out
						}
					case "R":
						target = pos + k
						if target > max {
							break out
						}
					}
					if g.grid[target] == "@" {
						alone=false
						break
					}
				}
			}
			if alone {
				count++
			}
		}
	}
	return count
}

func findNodesDestroyed(pos int, g Grid) int {
	var count, target int
	size := g.height * g.width
	min := pos - (pos % g.width)
	max := min + g.width - 1

	for _, d := range directions {
		out:
		for k := 1; k <= 3; k++ {
			switch d {
			case "U":
				target = pos - k*g.width
				if target < 0 {
					break out
				}
			case "D":
				target = pos + k*g.width
				if target >= size {
					break out
				}
			case "L":
				target = pos - k
				if target < min {
					break out
				}
			case "R":
				target = pos + k
				if target > max {
					break out
				}
			}

			if g.grid[target] == "#" {
				break
			} else if g.grid[target] == "@" {
				count++
			}
		}
	}
	return count
}

func FindBestPos(g Grid, rounds, bombs int) int {
	bestPos := -1
	maxNodesDestroyed := 0

	for i, v := range g.grid {
		if v == "@" || v == "#" {
			continue
		}
		nodesDestroyed := findNodesDestroyed(i, g)
		if nodesDestroyed > maxNodesDestroyed {
			ng := g.Copy()
			ng.Update(i, true)
			count := ng.getNbNodesAlone()
			if count >= bombs {
				continue
			}
			maxNodesDestroyed = nodesDestroyed
			bestPos = i
		}
	}
	return bestPos
}

func main() {
	var width, height int
	fmt.Scanln(&width, &height)

	var tmp string

	for i := 0; i < height; i++ {
		var mapRow string
		fmt.Scanln(&mapRow)
		tmp += mapRow
	}

	grid := strings.Split(tmp, "")
	g := Grid{width: width, height: height, grid: grid}
	nbNodesLeft := g.getNbNodesLeft()

	for {
		var rounds, bombs int
		fmt.Scanln(&rounds, &bombs)
		fmt.Fprintln(os.Stderr, g.grid)
		if nbNodesLeft > 0 {
			p := FindBestPos(g, rounds, bombs)
 			t := g.grid[p]
			g.Update(p, false)

			if t != "." {
				fmt.Println("WAIT")
				continue
			}
			nbNodesLeft = g.getNbNodesLeft()
			
			x := p % g.width
			y := p / g.width
			fmt.Printf("%d %d\n", x, y)
			continue
		}

		fmt.Println("WAIT")
	}
}
