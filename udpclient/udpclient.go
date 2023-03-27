package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 20000,
	})
	if err != nil {
		fmt.Printf("DialUDP failed,err=%v\n", err)
	}
	defer socket.Close()

	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" { //输入q则退出
			return
		}
		// 发送数据
		_, err = socket.Write([]byte(s))
		if err != nil {
			fmt.Printf("write failed,err=%v\n", err)
			return
		}
		// 接收数据
		data := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("ReadFromUDP failed,err=%v\n", err)
			return
		}

		fmt.Printf("来自[%v]的数据(%v):%v\n", remoteAddr, n, string(data[:n]))
	}
}
