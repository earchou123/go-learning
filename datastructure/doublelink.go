package main

import (
	"fmt"
)

type Node struct {
	no   int
	data string
	pre  *Node
	next *Node
}

func ListLink(head *Node) {
	tmpNode := head
	// 判断是否为空列表
	if tmpNode.next == nil {
		fmt.Println("Empty list")
		return
	}

	for {
		fmt.Printf("%v[%v] <=> ", tmpNode.next.no, tmpNode.next.data)
		tmpNode = tmpNode.next
		if tmpNode.next == nil {
			break
		}

	}
}

func ListLinkReverse(head *Node) {
	tmpNode := head
	// 判断是否为空列表
	if tmpNode.next == nil {
		fmt.Println("Empty list")
		return
	}
	// 使tmpNode为最后一个节点
	for {
		tmpNode = tmpNode.next
		if tmpNode.next == nil {
			break
		}
	}
	// 逆序遍历
	for {
		fmt.Printf("%v[%v] <=> ", tmpNode.no, tmpNode.data)
		tmpNode = tmpNode.pre
		if tmpNode.pre == nil {
			break
		}
	}
}

func InsertNode(head *Node, node *Node) {
	tmpNode := head

	for {
		if tmpNode.next == nil { //找到队尾，跳出循环
			break
		}
		tmpNode = tmpNode.next
	}
	tmpNode.next = node
	node.pre = tmpNode
}

func InsertNodeByNo(head *Node, node *Node) {
	tmpNode := head
	flag := true
	for {
		if tmpNode.next == nil { //找到队尾，跳出循环
			break
		} else if node.no == tmpNode.no { // 链中有该no，不插入
			flag = false
			break
		} else if node.no < tmpNode.next.no { // 插入到tmpNode后面
			break
		}
		tmpNode = tmpNode.next
	}
	if flag {
		node.next = tmpNode.next
		node.pre = tmpNode
		if tmpNode.next != nil { // tmpNode.next不为nil时，添加pre指针
			tmpNode.next.pre = node
		}
		tmpNode.next = node
	} else {
		fmt.Printf("node.no[%d] 已存在\n", node.no)
	}
}

func DelNode(head *Node, no int) {
	// 通过no删除节点
	tmpNode := head
	flag := false

	// 判断是否为空列表
	if tmpNode.next == nil {
		fmt.Println("Empty list")
		return
	}
	for {
		if tmpNode.next == nil { // 找到了队尾
			break
		} else if tmpNode.next.no == no { //找到了
			flag = true
			break
		}
		tmpNode = tmpNode.next
	}
	if flag {
		tmpNode.next = tmpNode.next.next
		if tmpNode.next != nil {
			tmpNode.next.pre = tmpNode
		}
	} else {
		fmt.Printf("没找到no[%v]\n", no)
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
	node4 := Node{
		no:   4,
		data: "node4",
	}
	node5 := Node{
		no:   3,
		data: "node5",
	}

	InsertNode(&node, &node1)
	InsertNode(&node, &node2)
	InsertNode(&node, &node4)
	ListLink(&node)
	fmt.Println("")

	InsertNodeByNo(&node, &node3)
	InsertNodeByNo(&node, &node5)
	ListLink(&node)
	fmt.Println("")

	ListLinkReverse(&node)
	fmt.Println("")

	// 删除节点
	DelNode(&node, 3)
	DelNode(&node, 999)
	ListLink(&node)

}
