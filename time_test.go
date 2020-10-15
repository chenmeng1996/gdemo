package main

import (
	"fmt"
	"testing"
	"time"
)

var (
	layout = "2006-01-02 15:04:05"
)

// 字符串 转 时间
func TestStr2Time(t *testing.T) {
	// 默认UTC时区
	tm, _ := time.Parse(layout, "2020-10-13 16:00:00")
	fmt.Println(tm, tm.Unix())

	// 指定上海时区
	tz, _ := time.LoadLocation("Asia/Shanghai")
	tm, _ = time.ParseInLocation(layout, "2020-10-13 16:00:00", tz)
	fmt.Println(tm, tm.Unix())

	// 当前操作系统所在时区
	tm, _ = time.ParseInLocation(layout, "2020-10-13 16:00:00", time.Local)
	fmt.Println(tm, tm.Unix())
}

// 时间转换时区
func TestChangeTimezone(t *testing.T) {
	tm := time.Now()
	fmt.Println(tm)

	// 当前操作系统时区
	tm1 := tm.Local()
	fmt.Println(tm1)

	// 上海时区
	tz, _ := time.LoadLocation("Asia/Shanghai")
	tm2 := tm.In(tz)
	fmt.Println(tm2)
}

// 时间 转 字符串
func TestTime2Str(t *testing.T) {
	tm := time.Now()
	str := tm.Format(layout)
	fmt.Println(str)
}

// 时间戳 转 时间
func TestUnix2Time(t *testing.T) {
	unix := time.Now().Unix()
	tm := time.Unix(unix, 0)
	fmt.Println(tm)
}

// 时间详细信息
func TestDetail(t *testing.T) {
	tm := time.Now()
	fmt.Println(tm.Year(), int(tm.Month()), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
}

// 时间加减
func TestChange(t *testing.T) {
	tm := time.Now()
	fmt.Println(tm)
	tm1 := tm.Add(time.Hour * 1) // 加时间段
	fmt.Println(tm1)
	tm2 := tm.Add(-time.Minute * 30) // 减时间段
	fmt.Println(tm2)

	tm3 := tm.AddDate(1, -1, 0) // 加减日期
	fmt.Println(tm3)
}

// 两个时间的差
func TestSub(t *testing.T) {
	tm := time.Now()
	tm2 := tm.AddDate(0, 0, 1)
	d := tm2.Sub(tm)
	fmt.Println(d, int(d))
}
