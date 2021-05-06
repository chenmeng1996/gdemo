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
