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