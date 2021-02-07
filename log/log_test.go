package log

import (
	"io"
	"log"
	"os"
	"testing"
)

//自定义logger
var (
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

//初始化log等级，输出方式，前缀，附带的信息
func init() {
	Debug = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

//默认日志输出到stderr
func TestLogPrint(t *testing.T) {
	log.Println("log println")
}

//输出日志后调用os.exit(1)
func TestLogFatal(t *testing.T) {
	log.Fatalln("log fatalln")
}

//输出日志后调用panic
func TestLogPanic(t *testing.T) {
	log.Panicln("log panicln")
}
func TestDebug(t *testing.T) {
	Debug.Println("log println") //[DEBUG] 2021/02/07 15:57:35 log_test.go:31: log println
}
func TestInfo(t *testing.T) {
	Info.Println("log fatalln") //[INFO] 2021/02/07 15:57:54 log_test.go:34: log fatalln
}
func TestError(t *testing.T) {
	Error.Println("log panicln") //[ERROR] 2021/02/07 15:58:00 log_test.go:37: log panicln
}

//输出日志到文件
func TestLogToFile(t *testing.T) {
	//打开文件读写，如果不存在则创建，如果存在则写入到末尾
	f, err := os.OpenFile("log_test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func() {
		//关闭文件
		err2 := f.Close()
		if err2 != nil {
			log.Fatalf("error closing file: %v", err2)
		}
	}()
	//设置输出到文件
	Info.SetOutput(f)
	Info.Println("This is a test log entry")
}

//同时输出日志到控制台和文件
func TestLogToConsoleAndFile(t *testing.T) {
	f, err := os.OpenFile("log_test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			log.Fatalf("error closing file: %v", err2)
		}
	}()
	//合并多个writer
	fileAndStdoutWriter := io.MultiWriter(f, os.Stdout)
	//设置输出到合并的writer
	Info.SetOutput(fileAndStdoutWriter)
	Info.Println("This is a test log entry")
}
