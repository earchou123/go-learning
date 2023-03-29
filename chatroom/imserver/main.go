package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Read fail err=%v\n", err)

	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Printf("json.Unmarshal err=%v\n", err)
		return
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		fmt.Println("读取客户端发送的数据。。。")
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了链接。")
				return
			} else {
				fmt.Printf("readPkg err=%v\n", err)
				return
			}
		}
		fmt.Printf("读到的mes=%v\n", mes)
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
