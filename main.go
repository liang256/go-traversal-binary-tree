package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

func buildTree() *Node {
	n1 := Node{nil, nil, 1}
	n2 := Node{nil, nil, 2}
	n3 := Node{nil, nil, 3}
	n4 := Node{nil, nil, 4}
	n5 := Node{nil, nil, 5}
	n6 := Node{nil, nil, 6}
	n7 := Node{nil, nil, 7}
	n6.right = &n7
	n5.right = &n6
	n4.right = &n5
	n4.left = &n2
	n2.left = &n1
	n2.right = &n3
	return &n4
}

func main() {
	root := buildTree()
	traversal_inorder(root)
}

func traversal_inorder(t *Node) {
	if t == nil {
		return
	}

	traversal_inorder(t.left)
	fmt.Println(t.value)
	traversal_inorder(t.right)
}
