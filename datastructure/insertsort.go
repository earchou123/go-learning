package main

import "fmt"

func InsertSort(arr []int) {
	// 从第二个元素开始循环，第一个元素看作有序
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]  // 第一个插入的元素
		insertIndex := i - 1 // 插入位置

		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 需要交换位置，元素后移
			insertIndex--                         // 插入坐标前移
		}
		// 插入
		if insertIndex+1 != i { // 如果index+1 = i，插入的数据已有序，不需要再插入
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("[%d] : %v \n", i, arr)
	}
}

func main() {
	arr := []int{4, 6, 2, 3, 0, 1, 5, 9}
	InsertSort(arr)
}
