package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ESQuery struct {
	Match ESMatch `json:"match"`
}
type ESMatch struct {
	ENV string `json:"env"`
}
type ESSort struct {
	Timestamp string `json:"@timestamp"`
}
type ESSearch struct {
	*ESQuery `json:"query,omitempty"`
	Sort     *ESSort `json:"sort"`
	Size     int     `json:"size"`
}
type Map = map[string]interface{}

var str string

func main() {
	n := strings.SplitN("asd.qwe.zcx.qwer", ".", 3)
	fmt.Println(n)
	esSearch := ESSearch{
		Sort: &ESSort{Timestamp: "desc"},
		Size: 0,
	}
	marshal, _ := json.Marshal(esSearch)
	fmt.Printf(string(marshal))
}
func init() {
	fmt.Println("json init")
}
