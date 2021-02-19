package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.IsAbs("/root"))                    //true
	fmt.Println(path.IsAbs("/"))                        //true
	fmt.Println(path.Split("/root/hello/world"))        ///root/hello/ world
	fmt.Println(path.Split("/"))                        ///
	fmt.Println(path.Join("root", "hello/", "/world/")) //root/hello/world
	fmt.Println(path.Dir("root/hello/"))                //root/hello
	fmt.Println(path.Dir(""))                           //.
	fmt.Println(path.Base("root/hello"))                //hello
	fmt.Println(path.Base("/"))                         ///
	fmt.Println(path.Base(""))                          //.
	fmt.Println(path.Ext("/a/b/c/bar.css"))             //.css

	fmt.Println(path.Clean("a/c"))                         //a/c
	fmt.Println(path.Clean("a//c"))                        //a/c
	fmt.Println(path.Clean("a/c/."))                       //a/c
	fmt.Println(path.Clean("a/c/b/.."))                    //a/c
	fmt.Println(path.Clean("/../a/c"))                     ///a/c
	fmt.Println(path.Clean("/../a/b/../././/c"))           ///a/c
	fmt.Println(path.Match("root/a*/*", "root/asdf/qwer")) //true <nil>
}
