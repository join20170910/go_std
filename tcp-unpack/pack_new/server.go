package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

const (
	BYTES_SIZE uint16 = 1024
	HEAD_SIZE  int    = 2
)

func StartServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		fmt.Println(conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doConn(conn)
	}
}
func doConn(conn net.Conn) {
	var (
		buffer           = bytes.NewBuffer(make([]byte, 0, BYTES_SIZE))
		bytes            = make([]byte, BYTES_SIZE)
		isHead      bool = true
		contentSize int
		head        = make([]byte, HEAD_SIZE)
		content     = make([]byte, BYTES_SIZE)
	)
	for {
		readLen, err := conn.Read(bytes)
		if err != nil {
			log.Println("Error reading", err.Error())
			return
		}
		_, err = buffer.Write(bytes[0:readLen])
		if err != nil {
			log.Println("Error writing to buffer", err.Error())
			return
		}

		for {
			if isHead {
				if buffer.Len() >= HEAD_SIZE {
					_, err := buffer.Read(head)
					if err != nil {
						fmt.Println("Error reading", err.Error())
						return
					}
					contentSize = int(binary.BigEndian.Uint16(head))
					isHead = false
				} else {
					break
				}
			}
			if !isHead {
				if buffer.Len() >= contentSize {
					_, err := buffer.Read(content[:contentSize])
					if err != nil {
						fmt.Println("Error reading", err.Error())
						return
					}
					fmt.Println(string(content[:contentSize]))
					isHead = true
				} else {
					break
				}
			}
		}
	}
}
