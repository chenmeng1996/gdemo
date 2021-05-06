package main

import (
	"fmt"
	"testing"
)

func TestInterface1(t *testing.T) {
	var a InterfaceA
	b := a
	fmt.Printf("%p, %p\n", a, b)   // 接口指向的内存地址，a和b都指向nil
	fmt.Printf("%p, %p\n", &a, &b) // 接口本身的内存地址，a和b不同

	c := A1{}
	a = &c
	c.a = 2
	fmt.Printf("%p\n", &c)

	fmt.Printf("%p, %p\n", a, b) // a指向指针c（存储的是c的内存地址），b还是指向nil
	a.Print()

	a = c
	c.a = 3
	fmt.Printf("%p, %p\n", a, b) // a指向结构体c的副本，b还是指向nil
	b = a
	a.Print()
	//b.Print() // 错误
}

type InterfaceA interface {
	Print()
}

type A1 struct {
	a int
}

func (a A1) Print() {
	fmt.Println(a)
}
