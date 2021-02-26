package main

import (
	"fmt"
	"testing"
)

type P struct {
	a int
}

// map get返回的是值拷贝
func TestMap1(t *testing.T) {
	m := make(map[string]P)
	m["a"] = P{1}

	x1 := m["a"]
	x2 := m["a"]
	fmt.Println(x1 == x2, &x1 == &x2)
}

func TestMap2(t *testing.T) {
	m := make(map[string]*P)
	m["a"] = &P{1}

	x1 := m["a"]
	x2 := m["a"]
	fmt.Println(x1 == x2, *x1 == *x2)
}

func TestMap3(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "123"

	delete(m, "a")
}

func TestMap4(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m)
}

func TestMap5(t *testing.T) {
	var m map[string]int
	fmt.Println(m["a"])
}

// 测试map之间的赋值
func TestMap6(t *testing.T) {
	m1 := map[string]int{"a": 1}
	m2 := m1
	fmt.Printf("%p %p %p %p\n", m1, m2, &m1, &m2)

	m2["a"] = 2
	fmt.Println(m1, m2)
}

func TestMap7(t *testing.T) {
	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"a": 2, "b": 3}
	for k, v := range m2 {
		m1[k] = v
	}
	fmt.Println(m1)
}
