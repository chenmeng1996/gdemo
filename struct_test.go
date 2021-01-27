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
