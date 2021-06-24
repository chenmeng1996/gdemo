package main

import (
	"fmt"
	"testing"
)

type C struct {
	A int
}

func TestCopy1(t *testing.T) {
	a := &C{A: 1}
	b := *a
	b.A = 2
	fmt.Println(a, b)
}
