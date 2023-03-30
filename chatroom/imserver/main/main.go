package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Printf("客户端和服务器的通讯协程错误 err=%v\n", err)
		return
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
