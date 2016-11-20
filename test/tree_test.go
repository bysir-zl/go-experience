package test

import (
	"testing"
	"log"
)

type Node struct {
	value int
	Child []*Node
}

func TestTree(t *testing.T) {
	root := &Node{value:0}
	node1 := &Node{value:1}
	node2 := &Node{value:2}
	node3 := &Node{value:3}
	root.Child = []*Node{node1, node2, node3}

	route := pathOne{}
	treeRoot(root, route)

	log.Printf("%v", path[2][1].value)

}

type pathOne []*Node
var path map[int]pathOne = map[int]pathOne{}

func treeRoot(root *Node, route pathOne) {
	route = append(route, root)
	if root.Child != nil {
		for _, v := range root.Child {
			treeRoot(v, route)
		}
	} else {
		if path[root.value] == nil {
			path[root.value] = pathOne{}
		}

		// u should copy the slice , then saved ;
		route3:=make(pathOne,len(route))
		copy(route3,route)
		path[root.value] = route3

		// if u not, u 
		//path[root.value] = route
	}
}
