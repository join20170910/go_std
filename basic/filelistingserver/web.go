package main

import (
	"awesomeProject/basic/filelistingserver/filelisting"
	"awesomeProject/go-tcp/lib/logger"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			logger.Warn("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.Handler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(any(err))
	}
}
