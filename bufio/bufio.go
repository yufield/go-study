package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	reader := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader := bufio.NewReader(reader)
	fmt.Println(newReader.Size()) //4096
	ls := make([]byte, 10)
	n, _ := newReader.Read(ls)
	fmt.Println(n)  //8
	fmt.Println(ls) //[97 98 99 100 101 102 103 104 0 0]//这里必须要注视掉然后下面才可以正常读取，因为不可多次读取
	reader2 := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader2 := bufio.NewReader(reader2)
	fmt.Println(newReader2.Size()) //4096
	lsls := make([]byte, 0, 15)    //注意，初始化的长度只能写0，否则会预存数据，达不到空长度的效果，但是我们的总容量必须要给一个，不给的话也会自动扩容的！但是效率不高。
	buffer := bytes.NewBuffer(lsls)
	fmt.Println("写入之前的buffer:", buffer)     //写入之前的buffer:
	fmt.Println(newReader2.WriteTo(buffer)) //这里不一定要给io.wirter接口的，我们给这个接口的实现也是可以的
	fmt.Println(lsls)                       //[]，有内容但长度是0，所以没有打印出来
	fmt.Println(lsls[:15])                  //[97 98 99 100 101 102 103 104 0 0 0 0 0 0 0]
	bb := buffer.Bytes()
	fmt.Println(bb) //[97 98 99 100 101 102 103 104]
	fmt.Printf("%p,%p\n", bb, lsls)
	// 使用make初始化切片
	//test := make([]int, 2, 5)
	// 对应的内存结构
	//test --> []int  -->	ptr(数组指针) | len(值为2) | cap(值为5)
	//ptr  --> [5]int -->	 0 | 0 | 0 | 0 | 0
	reader4 := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader4 := bufio.NewReader(reader4)
	fmt.Println(newReader4.Size()) //4096

	// Peek返回下一个n个字节，而不会使阅读器前进。 字节在下一个读取调用时停止有效。
	// 如果Peek返回的字节数少于n个字节，则它还会返回一个错误，说明读取短的原因。 如果n大于b的缓冲区大小，则错误为ErrBufferFull。
	//
	//调用Peek会阻止UnreadByte或UnreadRune调用成功，直到下一次读取操作为止。
	peek, _ := newReader4.Peek(2)
	fmt.Println(peek) //[97 98]
}
