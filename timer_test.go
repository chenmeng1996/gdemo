package main

import (
	"fmt"
	"testing"
	"time"
)

// 超时
func TestTimeout(t *testing.T) {
	ti := time.NewTimer(time.Second * 2)
	defer ti.Stop()

	resCh := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		resCh <- "hello"
	}()

	select {
	case <-ti.C:
		fmt.Println("超时")
	case s := <-resCh:
		fmt.Println(s)
	}
}

func TestTimeout2(t *testing.T) {
	ti := time.NewTimer(time.Second * 2)
	defer ti.Stop()

	resCh := make(chan string)
	//f := func() {
	//	time.Sleep(time.Second * 3)
	//	resCh <- "hello"
	//}()

	select {
	case <-ti.C:
		fmt.Println("超时")
	case s := <-resCh:
		fmt.Println(s)
	}
}

func TestTimeout3(t *testing.T) {
	stop := make(chan struct{})
	go func() {
		tick := time.Tick(4 * time.Second)
		select {
		case <-tick:
			stop <- struct{}{}
		}
	}()

	// 这是一个阻塞逻辑，最多阻塞4秒
	select {
	case <-stop:
		fmt.Println("退出")
	}
}
