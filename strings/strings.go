package main

import (
	"encoding/json"
	"fmt"
)

type Query struct {
	find   interface{}
	filter interface{}
}

func main() {
	//fmt.Println(strings.SplitN("aaaa.bbb.ccc", ".", 2))
	//fmt.Println(strings.SplitAfterN("aaaa.bbb.ccc", ".", 2))
	v := Query{
		find:   "asdf",
		filter: "asdfasdf",
	}
	marshal, _ := json.Marshal(v)
	fmt.Println(string(marshal))
}
