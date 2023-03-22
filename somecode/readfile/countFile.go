package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type charCount struct {
	ChCount    int // 记录字母个数
	NumCount   int // 数字个数
	SpaceCount int // 空格个数
	OtherCount int // 其他字符个数
	HanCount   int // 统计汉字个数
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	var count charCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if len(str) == 0 && err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Printf("read file err=%v\n", err)
		}

		runes := []rune(str)
		for _, v := range runes {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case unicode.Is(unicode.Scripts["Han"], v):
				count.HanCount++
			default:
				count.OtherCount++ // 包含了换行符
			}
		}
	}
	fmt.Printf("汉字：%v，字母：%v，数字：%v，空格%v，其他字符：%v\n", count.HanCount, count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
