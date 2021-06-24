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

func TestPoint2(t *testing.T) {
	type A struct {
		a int
	}
	var a1 = &A{}
	a2 := a1
	fmt.Println(a2)
	*a1 = A{a: 1}
	fmt.Println(a2)
}

func TestPoint3(t *testing.T) {
	a := func() {}
	b := func() {}
	c := b
	// 数据大小不确定，所以存储的是类型和指向另一内存地址
	// 引用地址
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", c)
	// 变量地址
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
	// 类型
	fmt.Printf("%T\n", a)  // 函数（类似指针）
	fmt.Printf("%T\n", &a) // 指向函数的指针
}

func TestPoint4(t *testing.T) {
	a := make(map[int]int)
	b := make(map[int]int)
	c := b

	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", c)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
	// 类型
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
}

func TestPoint5(t *testing.T) {
	a := 1
	b := 1
	c := b
	var d *int
	d = &b
	// 基本类型数据大小确定，所以a直接存储值，不存在引用其他内存

	// 编译错误，因为int不引用其他内存
	//fmt.Printf("%p\n", a)
	//fmt.Printf("%p\n", b)
	// 正确，因为*int引用其他内存地址
	fmt.Printf("%p\n", d)

	// 变量地址
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)
	fmt.Printf("%p\n", &d) //
	// 类型
	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%T\n", &a) // *int
}
