package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Tag struct {
	f1 int `f:"cm" d:"yy"`
}

func TestTag(t *testing.T) {
	tp := reflect.TypeOf(Tag{})
	sf, _ := tp.FieldByName("f1")
	// 字符串，包含所有KV
	fmt.Println(sf.Tag)
	// 从字符串中根据key查询value
	v, ok := sf.Tag.Lookup("f")
	fmt.Printf("%s, %t\n", v, ok)
}
