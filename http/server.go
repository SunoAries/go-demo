package main

import (
	"net"
	"fmt"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("requset", string(buf[:n]))
	}
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("conn err", err.Error())
			continue
		}
		fmt.Println("aaa")
		go handleConnection(conn)
	}
}
