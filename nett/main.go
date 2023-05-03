package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer func() {
		recover()
	}()
	fmt.Println("handleConnection defer g ")
	defer conn.Close()
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

// panic 终止 当前 协程的运行
// panic 在退出协程之前 会执行所有已注册的defer
// 在defer中执行recover,可拯救panic的协程

func main() {
	var err any = "异常报错"
	ln, err := net.Listen("tcp", "9999")

	defer func() {
		recover()
	}()
	if err != nil {

		panic(any("异常报错"))

	}
	for true {
		conn, err := ln.Accept()
		if err != nil {
			panic(any("异常报错"))
		}
		go handleConnection(conn)
	}
}
