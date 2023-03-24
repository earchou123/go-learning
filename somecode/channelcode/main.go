package main

import "fmt"

// 统计1-8000数字中，哪些是素数。开启多个gorotine，统计素数。
// 放入8000个整数到intChan中
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
		fmt.Printf("put: %v \n", i)
	}
	close(intChan)
}

// 读取intChan，将素数放入到primeChan中
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool // 是否为素数
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		// 判断是否为素数
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		// 将素数放入primeChan
		if flag {
			primeChan <- v
		}
	}
	exitChan <- true
}

func main() {
	var intChan = make(chan int, 100)
	var primeChan = make(chan int, 1000)
	var exitChan = make(chan bool, 4)

	go putNum(intChan)
	// 开启四个goroutine读取数据
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		// exitChan都读取完，则可以关闭primeChan
		close(primeChan)
	}()

	// 遍历primeChan
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}

		fmt.Printf("素数：%v \n", v)
	}
	fmt.Println("退出")
}
