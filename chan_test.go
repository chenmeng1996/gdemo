package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChan1(t *testing.T) {
	c := make(chan int)
	go func() {
		a := <-c
		fmt.Println(a)
	}()

	time.Sleep(2 * time.Second)
	c <- 1

	time.Sleep(1 * time.Second)
}

func TestChan2(t *testing.T) {
	c := make(chan int)
	go func() {
		// 读chan，chan为空时阻塞，chan不为空时读取，chan关闭时会读取默认值
		a := <-c
		fmt.Println(a)

		a, ok := <-c
		fmt.Println(a, ok)
	}()

	time.Sleep(2 * time.Second)
	close(c)

	time.Sleep(1 * time.Second)
}

func TestChan3(t *testing.T) {
	c := make(chan int)
	go func() {
		// range读取chan，当chan为空时阻塞，当chan关闭时直接跳过
		for a := range c {
			fmt.Println(a)
		}
	}()

	time.Sleep(2 * time.Second)
	close(c)

	time.Sleep(1 * time.Second)
}
