package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestOs(t *testing.T) {
	stat, _ := os.Stat("../gotest.txt")               //文件gotest.txt的信息
	fmt.Println(stat.Name())                          //文件名gotest.txt
	fmt.Println(stat.Size())                          //文件大小18字节,18B
	fmt.Println(stat.Mode())                          //文件模式位-rw-rw-rw-
	fmt.Println(stat.ModTime())                       //文件修改时间2021-02-05 14:11:26.8717006 +0800 CST
	fmt.Println(stat.IsDir())                         //文件是否是目录false
	fmt.Println(stat.Sys())                           //底层数据来源
	file, _ := os.Open("../gotest.txt")               //打开gotest.txt
	b := make([]byte, 4096)                           //申请4KB大小的内存
	read, _ := file.Read(b)                           //读取文件内容到[]byte,返回读取到的字节数
	fmt.Println(string(b))                            //this is gotest.txt
	fmt.Println(read)                                 //18
	path, _ := exec.LookPath("go")                    //在PATH中寻找可执行程序路径
	fmt.Println(path)                                 //C:\Go\bin\go.exe
	command := exec.Command(path, "version")          //调用path所在的程序参数为version
	output, _ := command.Output()                     //获取程序输出
	fmt.Print(string(output))                         //go version go1.15.7 windows/amd64
	fmt.Println(os.Args)                              //本程序运行参数
	hostname, _ := os.Hostname()                      //主机名
	fmt.Println(hostname)                             //DESKTOP-O3U7NVH
	fmt.Println(os.Environ())                         //系统环境变量
	getwd, _ := os.Getwd()                            //程序运行路径
	fmt.Println(getwd)                                //C:\Users\yuxin\go\src\awesomeProject\os
	fmt.Println(os.Getpagesize())                     //获取系统内存页的尺寸4096
	_ = os.Setenv("MyName", "Bob")                    //设置系统环境变量，出错会返回错误
	fmt.Println(os.Getenv("MyName"))                  //获取指定环境变量
	fmt.Println(os.Expand("I am $MyName", os.Getenv)) //I am Bob,把字符串中的${var}或$var替换成后面的映射函数所获取的内容
	fmt.Println(os.ExpandEnv("I am $MyName"))         //I am Bob,把字符串中的${var}或$var替换成环境变量,等于os.Expand("I am $MyName", os.Getenv)
	fmt.Println(os.Geteuid())                         //返回调用者的用户IDwindows-1,linux0
	fmt.Println(os.Getuid())                          //返回调用者的有效用户IDwindows-1,linux0
	fmt.Println(os.Getgid())                          //返回调用者的组IDwindows-1,linux0
	fmt.Println(os.Getegid())                         //返回调用者的有效组IDwindows-1,linux0
	groups, _ := os.Getgroups()                       //返回调用者所属的所有用户组的组ID,windows不支持
	fmt.Println(groups)                               //windows err 不支持,linux[0]
	fmt.Println(os.Getpid())                          //进程id 6672
	fmt.Println(os.Getppid())                         //父进程id 16568
	_ = os.Unsetenv("MyName")                         //清楚指定环境变量
	os.Clearenv()                                     //清楚所有环境变量
	os.Exit(0)                                        //退出程序0表示成功，其他表示失败
}
