package main

import (
	"log"
	"net/http"
	"time"
)

var Addr = ":1210"

func main() {
	//	1. 创建路由器
	//	2. 设备路由规则
	//	3. 创建服务器
	//	4. 监听端口并提供服务
	mux := http.NewServeMux()
	mux.HandleFunc("/bye", sayBye)
	server := &http.Server{
		Addr:         Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	log.Println("Starting http server at " + Addr)
	log.Fatal(server.ListenAndServe())
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("bye bye, this is httpServer"))

}
