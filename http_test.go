package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpResponseWithoutClose(t *testing.T) {
	for true {
		_, err := http.Get("https://www.baidu.com")
		if err != nil {
			fmt.Printf("error occurred while fetching page, error: %s", err.Error())
		}
		fmt.Println("ok")
	}
}

func TestHttpResponseWithClose(t *testing.T) {
	for true {
		resp, err := http.Get("https://www.baidu.com")
		if err != nil {
			fmt.Printf("error occurred while fetching page, error: %s", err.Error())
		}
		fmt.Println("ok")
		if resp != nil {
			resp.Body.Close()
		}
	}
}
