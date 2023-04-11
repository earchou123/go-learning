package main

import "fmt"

type Person struct {
	No   int
	Next *Person
}

// 添加节点
func AddPerson(num int) *Person { // 返回第一个节点
	first := &Person{}     // 第一个节点
	curPerson := &Person{} // 当前节点

	if num < 1 {
		fmt.Println("num less than 1")
		return first
	}

	for i := 1; i <= num; i++ {
		person := &Person{
			No: i,
		}
		// 第一个节点
		if i == 1 {
			first = person
			curPerson = person
			curPerson.Next = first
		} else {
			curPerson.Next = person
			curPerson = person
			curPerson.Next = first
		}
	}
	return first
}

func ShowPerson(first *Person) {
	// 判断空列表
	if first.Next == nil {
		fmt.Println("Empty link")
		return
	}

	curPreson := first
	for {
		fmt.Printf("person[%d] -> ", curPreson.No)
		if curPreson.Next == first {
			break
		}
		curPreson = curPreson.Next
	}
	fmt.Println("")
	return
}

/*
约瑟夫问题
设编号1,2,... n的人围坐一圈。约定编号为k(1<=k<=n)的人从1开始报数，数到m的人出列，它的下一位又从1开始报数，
数到m的人又出列，以此类推，知道所有人出列为止。由此禅师一个出队编号的序列
*/
func PlayGame(first *Person, startNo int, countNum int) {
	// 判断空列表
	if first.Next == nil {
		fmt.Println("Empty link")
		return
	}

	// 使tail指向最后一个元素
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	// first、tail节点移动startNo-1
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	// 循环删除
	for {
		// first、tail节点移动countNum-1
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("person[%d]出列\n", first.No)
		// 删除first节点
		first = first.Next
		tail.Next = first

		// tail == first 只剩一个人，退出循环
		if tail == first {
			fmt.Printf("person[%d]出列\n", first.No)
			break
		}
	}

}
func main() {
	first := AddPerson(20)
	ShowPerson(first)
	PlayGame(first, 5, 10)

}
