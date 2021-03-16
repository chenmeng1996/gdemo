package main

import (
	"fmt"
	"testing"
)

func TestInterface1(t *testing.T) {
	//var a InterfaceA
	//b := a
}

type InterfaceA interface {
	Print()
}

type A1 struct {
}

func (a A1) Print() {
	fmt.Println(123)
}
