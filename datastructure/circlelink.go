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
	if tmpNode.next == nil {
		fmt.Println("Empty list")
		return
	}

	for {
		fmt.Printf("%v[%v] => ", tmpNode.no, tmpNode.data)
		if tmpNode.next == head {
			break
		}
		tmpNode = tmpNode.next

	}
}
func InsertNode(head *Node, newNode *Node) {
	tmpNode := head
	// 添加第一个节点
	if tmpNode.next == nil {
		head.no = newNode.no
		head.data = newNode.data
		head.next = head
		return
	}

	for {
		if tmpNode.next == head { //找到队尾
			break
		}
		tmpNode = tmpNode.next
	}
	newNode.next = head
	tmpNode.next = newNode

}

func DelNode(head *Node, no int) *Node { // 删除需要返回一个head节点
	tmpNode := head
	// 判断空列表
	if tmpNode.next == nil {
		fmt.Println("Empty list")
		return head
	}
	// 仅有一个节点
	if tmpNode.next == head && tmpNode.no == no {
		tmpNode.next = nil
		return head
	}

	// 两个以上节点
	for {
		if tmpNode.next == head {
			if tmpNode.next.no == no { // 匹配到head节点
				// 删除head节点
				head = head.next // 更新head节点
				tmpNode.next = head
				return head
			}
			break
		}
		if tmpNode.next.no == no { // 匹配到节点
			// 删除节点
			tmpNode.next = tmpNode.next.next
			return head
		}
		tmpNode = tmpNode.next
	}
	fmt.Printf("未找到no[%d]\n", no)
	return head
}

func main() {
	node := &Node{}

	node1 := &Node{
		no:   1,
		data: "node1",
	}
	node2 := &Node{
		no:   2,
		data: "node2",
	}
	node3 := &Node{
		no:   3,
		data: "node3",
	}

	InsertNode(node, node1)
	InsertNode(node, node2)
	InsertNode(node, node3)
	ListLink(node)

	fmt.Println("")
	node = DelNode(node, 1)
	node = DelNode(node, 999)
	node = DelNode(node, 3)
	node = DelNode(node, 2)
	ListLink(node)
}
