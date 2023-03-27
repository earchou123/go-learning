package main

import (
	"fmt"
	"net"
)

func main() {
	// 建立UDP连接
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 20000,
	})
	if err != nil {
		fmt.Printf("listenUDP failed,err=%v\n", err)
	}

	defer listen.Close() // 关闭连接

	for {
		// 接收数据
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("readUDP failed,err=%v\n", err)
			continue
		}
		fmt.Printf("来自[%v]的消息：%v \n", addr, string(data[:n]))

		// 发送数据给客户端
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Printf("writeToUDP failed,err=%v\n", err)
			continue
		}

	}

}
