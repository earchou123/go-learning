package main

import "fmt"

func SelectSort(arr []int) {
	fmt.Printf("原始数组: %v\n", arr)
	for k := 0; k < len(arr); k++ {
		max := arr[k] // 最大值
		maxIndex := k // 最大值下标

		// 查找最大数，存储最大值、最大值下标
		for i := k + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		// 元素交换
		if maxIndex != k { // 最大值下表不等于k，即需要进行交换
			arr[k], arr[maxIndex] = arr[maxIndex], arr[k]
		}
		fmt.Printf("第%d次交换:%v\n", k+1, arr)
	}
}

func main() {
	arr := []int{4, 6, 2, 3, 0, 1, 5, 9}
	SelectSort(arr)
}
