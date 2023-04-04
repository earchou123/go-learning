package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (this *Queue) AddQueue(val int) (err error) {
	if this.rear == this.maxSize-1 {
		return errors.New("队列已满")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

func (this *Queue) ShowQueue() (err error) {
	if this.rear == this.front {
		return errors.New("队列为空")
	}
	fmt.Println("queue:")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%v\t", i, this.array[i])
	}
	fmt.Println()
	return
}

func (this *Queue) GetQueue() (val int, err error) {
	if this.rear == this.front {
		return -1, errors.New("队列为空")
	}
	this.front++
	val = this.array[this.front]
	return
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key int

	for {
		fmt.Println("1 AddQueue")
		fmt.Println("2 ShowQueue")
		fmt.Println("3 GetQueue")
		fmt.Println("4 exit")
		fmt.Println("input your choose(1-4):")
		fmt.Scanf("%v\n", &key)
		switch key {
		case 1:
			var val int
			fmt.Println("input val:")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err := queue.ShowQueue()
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Get val :%v \n", val)
			}
		case 4:
			os.Exit(0)
		}
	}
}
