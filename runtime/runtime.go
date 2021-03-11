package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go goroutine()
	fmt.Println(runtime.NumGoroutine()) //2
	fmt.Println(runtime.GOROOT())       //C:\Go
	fmt.Println(runtime.Version())      //go1.15.7
	fmt.Println(runtime.NumCPU())       //6
}
func goroutine() {
	time.Sleep(time.Second)
}
