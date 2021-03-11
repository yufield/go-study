package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	timeq, _ := time.Parse("2006-01-02T15:04:05.000Z", "2021-02-22T09:57:19.237Z")
	fmt.Println(timeq)
	//time1, _ := time.Parse("2006-01-02 15:04:05 MST", "2021-02-05 14:27:20 CST")
	//time2 := time.Date(2009, time.November, 10, 23, 2, 3, 6, time.UTC)
	//format := time1.Format("2006年1月2日15点4分5秒")
	//fmt.Println(time.Monday)                //Monday
	//fmt.Println(time.January)               //January
	//fmt.Println(time1)                      //2021-02-05 14:27:20 +0800 CST
	//fmt.Println(format)                     //2021年2月5日14点27分20秒
	//fmt.Println(time1.Year())               //2021
	//fmt.Println(time1.Month())              //February
	//fmt.Println(time1.Day())                //5
	//fmt.Println(time1.Hour())               //14
	//fmt.Println(time1.Minute())             //27
	//fmt.Println(time1.Second())             //20
	//fmt.Println(time1.Date())               //2021 February 5
	//fmt.Println(time1.UTC())                //2021-02-05 06:27:20 +0000 UTC
	//fmt.Println(time1.Zone())               //CST 28800
	//fmt.Println(time1.Equal(time1))         //true
	//fmt.Println(time1.Equal(time1.UTC()))   //true
	//time.Now()                              //本地时区时间
	//time.Sleep(100 * time.Millisecond)      //当前goroutine睡眠100ms
	//fmt.Println(time2.Local())              //2009-11-11 07:02:03.000000006 +0800 CST
	//fmt.Println(time2)                      //2009-11-10 23:02:03.000000006 +0000 UTC
	//fmt.Println(time1.Equal(time2.Local())) //false
	//fmt.Println(time1.Equal(time2))         //false
	//fmt.Println(time2.Equal(time2.Local())) //true
	//fmt.Println(time1.IsZero())             //false
	//fmt.Println(time1.Location())           //Local
	//fmt.Println(time1.UTC().Location())     //UTC
	//fmt.Println(time1.Location().String())  //Local
	//fmt.Println(time1.Unix())               //1612506440
	//fmt.Println(time1.UnixNano())           //1612506440000000000
	//var Dur = time.Hour                     //1小时类型time.Duration
	//fmt.Println(int64(Dur))                 //转换成整形为3600000000000
	//fmt.Println(Dur.String())               //1h0m0s
	//fmt.Println(Dur.Hours())                //1
	//fmt.Println(Dur.Minutes())              //60
	//fmt.Println(Dur.Seconds())              //3600
	//fmt.Println(Dur.Milliseconds())         //3600000(就是3600e3)
	//fmt.Println(Dur.Microseconds())         //3600000000(就是3600e6)
	//fmt.Println(Dur.Nanoseconds())          //3600000000000(就是3600e9)
}
