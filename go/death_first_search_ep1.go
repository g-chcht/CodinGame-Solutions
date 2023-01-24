package main

import "fmt"

type Node struct {
    id int
    neighbours []*Node
	parent *Node
	isGateway, visited bool
}

func NewNode(id int) *Node {
    neighbours := make([]*Node, 0)
    d := Node{id: id, neighbours: neighbours}
    return &d
}

func (node *Node) AddNeighbour(i int, nodes []*Node){
	node.neighbours = append(node.neighbours, nodes[i])
}

func BFS(root *Node) (int,int) {
	c1, c2 := -1,-1
	queue := make([]*Node,0)
	
	root.visited = true
	queue = append(queue, root)
	s := 1

	for s > 0 {
		node := queue[0]
		s--
		queue = queue[1:]
		if node.isGateway {
			c1 = node.parent.id
			c2 = node.id
			break
		}
		for _,v := range node.neighbours {
			if !v.visited {
				v.visited = true
				v.parent = node
				queue = append(queue, v)
				s++
			}
		}
	}

	return c1, c2
}

func main() {
    // N: the total number of nodes in the level, including the gateways
    // L: the number of links
    // E: the number of exit gateways
    var N, L, E int
    fmt.Scan(&N, &L, &E)

	nodes := make([]*Node, N)
	for i:=0;i<N;i++{
		nodes[i] = NewNode(i)
	}
    
    for i := 0; i < L; i++ {
        // N1: N1 and N2 defines a link between these nodes
        var N1, N2 int
        fmt.Scan(&N1, &N2)
		nodes[N1].AddNeighbour(N2, nodes)
		nodes[N2].AddNeighbour(N1, nodes)
    }
    for i := 0; i < E; i++ {
        // EI: the index of a gateway node
        var EI int
        fmt.Scan(&EI)
		nodes[EI].isGateway = true
    }
    for {
        // SI: The index of the node on which the Bobnet agent is positioned this turn
        var SI int
        fmt.Scan(&SI)

		bobnet := nodes[SI]

		c1, c2 := BFS(bobnet)


        
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        // Example: 0 1 are the indices of the nodes you wish to sever the link between
        fmt.Println(c1, c2)
		for _, v := range nodes {
			v.parent = nil
			v.visited = false
		}
    }
}