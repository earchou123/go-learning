package main

import "fmt"

type Node struct {
	data string
	next *Node
}

func ListLink(head *Node) {
	tmpNode := head
	for {
		fmt.Printf("%v next:%v\n", tmpNode.data, tmpNode.next)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}
}

func InsertNode(head *Node, node *Node) {
	tmpNode := head
	for {
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}
	tmpNode.next = node
}

func main() {

	node1 := Node{
		data: "node1",
	}
	node2 := Node{
		data: "node2",
	}
	node3 := Node{
		data: "node3",
	}
	InsertNode(&node1, &node2)
	InsertNode(&node1, &node3)
	ListLink(&node1)

}
