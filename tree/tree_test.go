package tree

import (
	"testing"
	"github.com/deepzz0/go-com/log"
)

type Node struct {
	value int
	Child []Node
}

func TestTree(t *testing.T) {
	root := Node{value:0}
	node1 := Node{value:1}
	node2 := Node{value:2}
	node3 := Node{value:3}
	root.Child = []Node{node1, node2, node3}
	route := []*Node{}

	p:=P{}
	p.root = root
	p.path = map[int][]*Node{}
	p.treeRoot(p.root,route)
	log.Printf("%v", p.path[1][1].value)

	p2:=P2{}
	p2.root = root
	p2.path = map[int][]*Node{}
	p2.treeRoot(&p2.root,route)
	log.Printf("%v", p2.path[1][1].value)
}

type P struct {
	path  map[int][]*Node
	root Node
}
func ( p *P)treeRoot(root Node, route []*Node) {
	route = append(route, &root)
	if root.Child != nil {
		for _, v := range root.Child {
			p.treeRoot(v, route)
		}
	} else {
		p.path[root.value] = route
	}
}


type P2 struct {
	path  map[int][]*Node
	root Node
}
func ( p *P2)treeRoot(root *Node, route []*Node) {
	route = append(route, root)
	if root.Child != nil {
		for i, _ := range root.Child {
			p.treeRoot(&root.Child[i], route) //v
		}
	} else {
		p.path[root.value] = route
	}
}
