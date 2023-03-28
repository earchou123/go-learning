package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 8096)

		fmt.Println("读取客户端发送的数据。。。")

		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Printf("读到的buf=%v\n", buf[:4])
	}
}

func main() {
	fmt.Println("服务器在监听20000端口")
	listen, err := net.Listen("tcp", "0.0.0.0:20000")
	defer listen.Close()
	if err != nil {
		fmt.Printf("net.Listen err=%v\n", err)
		return
	}
	for {
		fmt.Println("等待客户端连接服务器..")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("listen.Accept err=%v\n", err)
		}
		go process(conn)
	}

}
