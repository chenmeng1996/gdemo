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

	notifyCh := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		notifyCh <- "hello"
	}()

	select {
	case <-ti.C:
		fmt.Println("超时")
	case s := <-notifyCh:
		fmt.Println(s)
	}
}
