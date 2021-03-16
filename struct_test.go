package main

import (
	"encoding/json"
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

type A struct {
	A int `json:"a"`
	B
}

type B struct {
	b string
}

func TestStructTwo(t *testing.T) {
	a := A{A: 1}
	fmt.Println(a)
	b := B{b: "qwe"}
	a.B = b

	res, _ := json.Marshal(a)
	fmt.Println(string(res))
}

func TestStructEqual(t *testing.T) {
	type E struct {
		A int
	}
	a := E{A: 1}
	b := E{A: 1}
	fmt.Println(a == b)
}
