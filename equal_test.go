package main

import (
	"fmt"
	"testing"
)

type Equal struct {
	a int
	b []int
}

func TestEqual1(t *testing.T) {
	e1 := &Equal{
		a: 1,
		b: []int{2, 3},
	}
	e2 := e1
	fmt.Println(e1 == e2)

	m := make(map[*Equal]int)
	m[e1] = 1
	m[e2] = 2
	fmt.Println(m)
}
