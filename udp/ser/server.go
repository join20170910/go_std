package main

import (
	"fmt"
	"net"
)

func main() {

	// 1、创建 监听服务器
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})

	if err != nil {
		fmt.Println("listen failed, err:%v\n", err)
		return
	}
	// 2、循环读取消息内容
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read failed from addr: %v, err:%v\n",
				addr, err)
			break
		}
		go func() {
			// todo sth 业务逻辑
			//3、 回复数据
			fmt.Printf("addr: %v data: %v count: %v\n", addr, string(data[:n]), n)
			_, err = listen.WriteToUDP([]byte("received success!"), addr)
			if err != nil {
				fmt.Printf("write failed, err: %v\n", err)
			}
		}()
	}

}
