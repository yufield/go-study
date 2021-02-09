package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello 你好"
	fmt.Println(len(str))                    //12
	fmt.Println(len([]rune(str)))            //8
	fmt.Println(utf8.RuneCountInString(str)) //8
}
