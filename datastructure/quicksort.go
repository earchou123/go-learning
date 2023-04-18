package main

import (
	"fmt"
)

func partition(left int, right int, arr []int) int {
	pivot := arr[(left+right)/2]
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	fmt.Println(arr)
	return left
}

func QuickSort(left int, right int, arr []int) {
	if left >= right {
		return
	}
	p := partition(left, right, arr)
	QuickSort(left, p-1, arr)
	QuickSort(p, right, arr)
}

func main() {
	arr := []int{4, 0, 2, 3, 6, 1, 5, 9, 8}
	fmt.Println("初始数组", arr)
	QuickSort(0, len(arr)-1, arr)
	fmt.Println("最终数组", arr)
}
