package main

import (
	"io"
	"log"
	"net/http"
)

var s = http.Server{
	Addr:              "127.0.0.1:8000", //地址端口
	Handler:           nil,
	TLSConfig:         nil,
	ReadTimeout:       0, //读请求超时时间
	ReadHeaderTimeout: 0, //读请求头超时时间
	WriteTimeout:      0, //写响应超时时间
	IdleTimeout:       0, //keep-alives打开时请求间的等待时间
	MaxHeaderBytes:    0, //请求头最大字节数
	TLSNextProto:      nil,
	ConnState:         nil,
	ErrorLog:          nil, //指定logger
	BaseContext:       nil,
	ConnContext:       nil,
}

func main() {
	http.HandleFunc("/health-check", HealthCheckHandler)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
