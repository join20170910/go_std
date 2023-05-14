package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type HandlerNewFunc func(response http.ResponseWriter, r *http.Request)

func (f HandlerNewFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)

}
func main() {
	hf := HandlerNewFunc(HelloHandler)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", bytes.NewBuffer([]byte("test")))
	hf.ServeHTTP(resp, req)
	bts, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bts))

}

func HelloHandler(response http.ResponseWriter, r *http.Request) {
	response.Write([]byte("hello world"))
}
