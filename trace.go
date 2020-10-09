package main

import (
	"log"
	"os"
	"runtime"
	"time"
)

func Trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	_, _ = os.Stdout.Write(buf[:n])
}
