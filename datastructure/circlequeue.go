package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	arrary  [5]int
	head    int
	tail    int
}

// 入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.isFull() {
		return errors.New("queue is full")
	}
	this.arrary[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// 出队列
func (this *CircleQueue) Pop() (pop int, err error) {
	if this.isEmpty() {
		return -1, errors.New("queue is empty")
	}
	//
	pop = this.arrary[this.head]
	this.head++
	return
}

// 显示队列
func (this *CircleQueue) ListQueue() (err error) {
	if this.isEmpty() {
		return errors.New("queue is empty")
	}
	fmt.Println("------打印环形队列-----")
	tmpHead := this.head
	size := this.Size()
	fmt.Println("size", size)
	for i := this.head; i < size; i++ {
		fmt.Println(this.arrary[tmpHead])
		tmpHead = (tmpHead + 1) % this.maxSize
	}
	fmt.Println("----------end----------")
	return
}

func (this *CircleQueue) isFull() bool {
	return this.head == (this.tail+1)%this.maxSize
}

func (this *CircleQueue) isEmpty() bool {
	return this.head == this.tail
}

func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	var circleQueue = CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key int
	for {
		fmt.Println("1 PushQueue")
		fmt.Println("2 ListQueue")
		fmt.Println("3 PopQueue")
		fmt.Println("4 exit")
		fmt.Println("input your choose(1-4):")
		fmt.Scanf("%v\n", &key)
		var val int
		switch key {
		case 1:
			fmt.Println("input val")
			fmt.Scanf("%v\n", &val)
			err := circleQueue.Push(val)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err := circleQueue.ListQueue()
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			pop, err := circleQueue.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Pop=", pop)
			}
		case 4:
			os.Exit(0)
		}
	}
}
