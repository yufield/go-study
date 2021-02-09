package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	// 构造一个listener
	listener, _ := net.Listen("tcp", "127.0.0.1:6379")
	defer func() { _ = listener.Close() }()
	for {
		// 接收请求
		conn, _ := listener.Accept()
		fmt.Println("conn")
		// 连接的处理
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	//defer conn.Close()
	//reader := bufio.NewReader
	var respBody = "<h1>Hello World<h1>"
	i := len(respBody)
	//返回的http消息头
	var respHeader = "HTTP/1.1 200 OK\n" +
		"Content-Type: text/html;charset=ISO-8859-1\n" +
		"Content-Length: " + strconv.Itoa(i)
	resp := respHeader + "\n\r\n" + respBody
	fmt.Println(resp)
	conn.Write([]byte(resp))
	//for true {
	//	line, _, err := reader.ReadLine()
	//	if err != nil {
	//		fmt.Println("err:", err)
	//		return
	//	}
	//	fmt.Println(string(line))
	//}
}
