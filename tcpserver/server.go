package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 处理完关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("reader failed, err=%v\n", err)
		}
		recv := string(buf[:n])
		fmt.Println("接收到的数据：", recv)
		conn.Write([]byte("ok")) // 回复客户端ok

	}
}

func main() {
	// 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Printf("listen failed,err=%v\n", err)
		return
	}
	for {
		// 等待客户端建立链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept fialed,err=%v\n", err)
			continue
		}
		// 启动一个gorotine处理连接
		go process(conn)
	}

}
