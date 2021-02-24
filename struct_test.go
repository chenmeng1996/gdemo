package main

import (
	"fmt"
	"testing"
)

type People struct {
	int
	string
}

func TestStruct1(t *testing.T) {
	p := People{1, "a"}
	fmt.Println(p.int, p.string)
}

func TestStruct2(t *testing.T) {
	p1 := People{1, "a"}
	p2 := People{2, "b"}
	p1 = p2
	fmt.Println(p1, p2)
	fmt.Printf("%p,%p\n", &p1, &p2)
	p2.string = "c"
	fmt.Println(p1, p2)
}

func TestStructInit(t *testing.T) {
	type A struct {
		Slice []string
		Map   map[string]int
	}
	a := A{
		Map: make(map[string]int),
	}
	fmt.Println(a, a.Slice == nil, a.Map == nil)
	a.Slice = append(a.Slice, "abc")
	a.Map["x"] = 1
	fmt.Println(a)
}
