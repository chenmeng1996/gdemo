package main

import (
	"fmt"
	"testing"
)

func TestPointAddress(t *testing.T) {
	a := 1
	b := &a
	fmt.Printf("%p %p", b, &b) // 指针指向的实际存储值的内存地址 指针本身的内存地址
	fmt.Printf("%p", &a)       // 值的地址 = 指针指向的实际存储值的内存地址
}
