package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 客户端连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Printf("Dial faile err=%v\n", err)
		return
	}

	input := bufio.NewReader(os.Stdin)
	for {
		// 接收输入内容
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s) // 去除空格
		// 输入q退出
		if strings.ToUpper(s) == "Q" {
			return
		}

		// 给服务端发送消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed,err=%v\n", err)
			return
		}

		// 接收服务端回复内容
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed,err=%v\n", err)
			return
		}
		fmt.Println("收到服务端回复：", string(buf[:n]))

	}

}
