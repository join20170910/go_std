package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, // 长连接超时时间
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  // 100-continue 状态码超时时间
	}
	//创建客户端
	client := &http.Client{
		Timeout:   time.Second * 30, //请求超时时间
		Transport: transport,
	}
	// 请求数据
	resp, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		fmt.Printf("请求错误 %v\n", err.Error())
	}
	defer resp.Body.Close()
	// 读取内容
	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("请求错误 %v\n", err.Error())
	}
	fmt.Println(string(bds))

}
