package main

import (
	"fmt"
)

type Node struct {
    digit byte
    nextDigits []*Node 
}

func NewNode(digit byte) *Node {
    nextDigits := make([]*Node, 0)
    d := Node{digit: digit, nextDigits: nextDigits}
    return &d
}

func (n *Node) AddToNode(digit byte) (*Node, int) {
    for _, node := range n.nextDigits {
        if node.digit == digit {
            return node, 0
        }
    }
    node := NewNode(digit)
    n.nextDigits = append(n.nextDigits, node)
    return node, 1
}

func main() {
    var N int
    fmt.Scan(&N)

    root := NewNode('R')
    currentNode := root
    sum, s := 0,0

    for i := 0; i < N; i++ {
        var telephone string
        fmt.Scan(&telephone)
        for _, digit := range telephone {
            currentNode, s = currentNode.AddToNode(byte(digit))
            sum += s
        }
        currentNode = root
    }

    fmt.Println(sum)
}