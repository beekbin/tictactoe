package tree

import (
	//"github.com/golang/glog"
	"fmt"
)

type Tree struct {
	root *Node
}

func NewTree(n *Node) *Tree {
	return &Tree{
		root: n,
	}
}

func (t *Tree) GetRoot() *Node {
	return t.root
}

func (t *Tree) SetRoot(n *Node) {
	t.root = n
}

func (t *Tree) AddChild(n *Node) {
	t.root.Children = append(t.root.Children, n)
}

func (t *Tree) SelectLeafNode() *Node {
	node := t.root
	node.visitCount += 1

	i := 0
	for len(node.Children) > 0 {
		node = node.SelectMostUCTChild()
		i ++
	}

	//if i > 7 {
	//	fmt.Printf("tree depth = %d\n", i)
	//}

	return node
}

func (t *Tree) PrintTree() {
	node := t.root

	fmt.Printf("------------------\n")
	fmt.Printf("%v\n", node)
	node.board.PrintBoard()

	if len(node.Children) > 0 {
		for _, child := range node.Children {
			fmt.Printf("\t%v\n", child)
			child.board.PrintBoard()
		}
	}
}