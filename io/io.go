package main

import (
	"fmt"
	"io"
	"os"
)

type alphaReader struct {
	// alphaReader 里组合了标准库的 io.Reader
	reader io.Reader
}

func (r alphaReader) Read(b []byte) (n int, err error) {
	i := 0
	for ; i < len(b); i++ {
		b[i] = 'A'
	}
	return i, nil
}
func main() {
	//reader := strings.NewReader("Clear is better than clever")
	reader := alphaReader{}
	p := make([]byte, 4)

	for i := 0; i < 50; i++ {
		n, err := reader.Read(p) //n返回读取的字节数
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
}
func init() {
	fmt.Println("io init")
}
