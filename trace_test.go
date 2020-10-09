package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTrace(t *testing.T) {
	defer Trace("operation")()
	time.Sleep(2 * time.Second)
}

func TestPrintStack(t *testing.T) {
	defer PrintStack()
	fmt.Print(1)
}
