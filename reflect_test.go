package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect1(t *testing.T) {
	s := "123"
	tp := reflect.TypeOf(s)
	fmt.Println(tp.Name(), tp.Kind().String())
}

func TestReflect2(t *testing.T) {
	var a []int
	fmt.Println(reflect.TypeOf(nil) == nil, reflect.TypeOf(a) == nil, a == nil)
}
