package nest

import "fmt"

type NestHuman struct {
	school string
}

func (n *NestHuman) SayHi() {
	fmt.Println("hello this is NestHuman")
}

func init() {
	fmt.Println("hello this is nest package")
}
