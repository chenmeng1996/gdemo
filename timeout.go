package main

import (
	"fmt"
	"time"
)

// 超时
func main() {
	t := time.NewTimer(time.Second * 2)
	defer t.Stop()

	notifyCh := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		notifyCh <- "hello"
	}()

	select {
	case <-t.C:
		fmt.Println("超时")
	case s := <-notifyCh:
		fmt.Println(s)
	}
}
