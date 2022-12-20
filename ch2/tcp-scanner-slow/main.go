package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// 端口已关闭或已过滤
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
