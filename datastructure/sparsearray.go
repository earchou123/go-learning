package main

import "fmt"

type ValNode struct {
	row int // 行
	col int // 列
	val int // 值
}

func main() {
	// 创建一个棋盘
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 白子

	fmt.Println("------原始棋盘------")
	for _, v := range chessMap {
		fmt.Println(v)
	}

	// 定义一个稀疏数组
	var sparseArr []ValNode
	// 第一个节点，记录二维数组的规模和默认值
	valNode := ValNode{
		row: 11, // 行
		col: 11, // 列
		val: 0,  // 默认值
	}
	sparseArr = append(sparseArr, valNode)
	// 转成稀疏数组
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}

	}
	fmt.Println("------稀疏数组------")
	for i, v := range sparseArr {
		fmt.Printf("%d : %d %d %v \n", i, v.row, v.col, v.val)
	}
}
