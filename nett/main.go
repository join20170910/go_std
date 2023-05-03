package main

import (
	"fmt"
	"net"
)

func main() {

	ln, err1 := net.Listen("tcp", "9999")
	if err1 != nil {
		return
	}
	conn, err := ln.Accept()
	if err != nil {
		panic(any("异常报错"))
	}
	var body [1024]byte
	for true {
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("收到消息: %s\n", string(body[:]))
		_, err = conn.Write(body[:])

		if err != nil {
			break
		}
	}
}
