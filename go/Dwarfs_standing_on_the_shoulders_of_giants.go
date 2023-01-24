package main

import (
	"fmt"
)

type Node struct {
    id int
    neighbours []*Node
}

func NewNode(id int) *Node {
    neighbours := make([]*Node, 0)
    d := Node{id: id, neighbours: neighbours}
    return &d
}

func DFS(node *Node, depth int) int {
	a := depth
	for _, v := range node.neighbours {
		r := DFS(v, depth+1)
		if r > a {
			a = r
		}
	}
	return a
}

func main() {
    // n: the number of relationships of influence
    var n int
    fmt.Scan(&n)

	nodes := make(map[int]*Node)
    
    for i := 0; i < n; i++ {
        // x: a relationship of influence between two people (x influences y)
        var x, y int
        fmt.Scan(&x, &y)
		n1, f1 := nodes[x]
		n2, f2 := nodes[y]

		if !f1 {
			n1 = NewNode(x)
			nodes[x] = n1
		}
		if !f2 {
			n2 = NewNode(y)
			nodes[y] = n2
		}
		n1.neighbours = append(n1.neighbours, n2)
    }

	maxDepth := 0
	for _,v := range nodes {
		d := DFS(v,1)
		if d > maxDepth {
			maxDepth = d
		}
	}

    fmt.Println(maxDepth)
}