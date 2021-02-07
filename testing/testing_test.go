package testing

import (
	"log"
	"testing"
)

//测试运行前初始化和测试运行后清理
func TestMain(m *testing.M) {
	log.Println("begin")
	m.Run() //测试运行
	log.Println("end")
}

//功能测试testing.T
//go test go-study/testing -test.run TestLog -v
func TestLog(t *testing.T) {
	t.Log("function testing")
}

//FailNow终止测试
func TestFailNow(t *testing.T) {
	t.FailNow()
}
func TestFail(t *testing.T) {
	t.Log("first fail")
	t.Fail()
	t.Log("second fail")
	t.Fail()
	t.Log("after fail")
}

//压力测试testing.B
//go test -v -bench=B ./testing/testing_test.go -test.run BenchmarkB
//打印内存分配-benchmem
func BenchmarkB(b *testing.B) {
	var n int
	// 重置计时器
	b.ResetTimer()
	// 停止计时器
	b.StopTimer()
	// 开始计时器
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n++
	}
	b.Log("total ", n)
}
