package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go goroutine()
	fmt.Println(runtime.NumGoroutine()) //2
}
func goroutine() {
	time.Sleep(time.Second)
}
