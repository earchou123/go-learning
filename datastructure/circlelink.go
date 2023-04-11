package main

import (
	"fmt"
)

type Node struct {
	no   int
	data string
	next *Node
}

func ListLink(head *Node) {
	tmpNode := head

	// 判断列表为空
	if tmpNode.next == head || tmpNode.next == nil {
		fmt.Println("l Empty list")
		return
	}

	for {
		fmt.Printf("%v[%v] => ", tmpNode.next.no, tmpNode.next.data)
		tmpNode = tmpNode.next
		if tmpNode.next == head {
			break
		}

	}
}
func InsertNode(head *Node, newNode *Node) {
	tmpNode := head

	for {
		if tmpNode.next == head || tmpNode.next == nil { //找到队尾
			break
		}
		tmpNode = tmpNode.next
	}
	newNode.next = head
	tmpNode.next = newNode

}

func DelNode(head *Node, no int) {
	tmpNode := head
	// 判断空列表
	if tmpNode.next == head || tmpNode.next == nil {
		fmt.Println("Empty list")
		return
	}

	// 仅有一个节点
	if tmpNode.next == head && tmpNode.no == no {
		tmpNode.next = nil
		return
	}

	flag := false
	for {
		if tmpNode.next == head { //找到队尾
			break
		} else if tmpNode.next.no == no {
			flag = true
			break
		}
		tmpNode = tmpNode.next
	}
	if flag {
		tmpNode.next = tmpNode.next.next
	} else {
		fmt.Printf("未找到no[%d]\n", no)
	}
}

func main() {
	node := Node{}

	node1 := Node{
		no:   1,
		data: "node1",
	}
	node2 := Node{
		no:   2,
		data: "node2",
	}
	node3 := Node{
		no:   3,
		data: "node3",
	}

	InsertNode(&node, &node1)
	InsertNode(&node, &node2)
	InsertNode(&node, &node3)
	ListLink(&node)

	fmt.Println("")
	DelNode(&node, 2)
	DelNode(&node, 999)
	ListLink(&node)
}
