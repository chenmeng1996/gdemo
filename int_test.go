package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInt10To2(t *testing.T) {
	a := int64(10)
	b := strconv.FormatInt(a, 16)
	fmt.Println(b)
}
