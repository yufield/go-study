package main

import (
	"flag"
	"fmt"
)

//设置flag名称和默认值，返回*int
var port = flag.Int("port", 80, "help message for port")
var open = flag.Bool("bool", false, "boolean")

//也可以申明变量后将地址传给flag.StringVar
var name string

func main() {
	fmt.Println(*port, name)
	// go run ./flag/flag.go
	// 80 alice
	// go run ./flag/flag.go -name bob --port 1234
	// 1234 bob
	// go run ./flag/flag.go -name bob carol
	// 80 bob
	fmt.Println(*open)
	//go run ./flag/flag.go
	//false
	//go run ./flag/flag.go -bool
	//true
	//go run ./flag/flag.go --bool
	//true
	//go run ./flag/flag.go -bool=false
	//false

}
func init() {
	//将地址传给flag.StringVar
	flag.StringVar(&name, "name", "alice", "help message for name")
	flag.Parse()
	fmt.Println("flag parsed")
}
