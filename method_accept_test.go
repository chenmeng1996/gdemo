package main

import (
	"fmt"
	"testing"
)

// 展示了 值接收者/指针接收者的区别
func Test1(t *testing.T) {
	p1 := Person{}
	fmt.Printf("address is %p\n", &p1)
	p1.Get()
	p1.Set()

	fmt.Println()

	p2 := &Person{}
	fmt.Printf("address is %p\n", p2)
	p2.Get()
	p2.Set()

	fmt.Println()

	// 只要结构体有指针接收者方法，接口只能赋值指针
	var age1 Age = p2
	fmt.Printf("address is %p\n", age1)
	age1.Get()
	age1.Set()

	// 编译不通过
	//var _ Age = p1

	fmt.Println()
}

type Person struct {
	age int
}

func (p Person) Get() {
	fmt.Printf("address is %p\n", &p)
}

func (p *Person) Set() {
	fmt.Printf("address is %p\n", p)
}

type Age interface {
	Get()
	Set()
}
