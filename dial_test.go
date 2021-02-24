package main

import (
	"fmt"
	"net"
	"testing"
)

func TestDial1(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:9090")
	fmt.Println(conn, err)
}
