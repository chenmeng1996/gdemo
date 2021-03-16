package main

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	f := func() error {
		var err error
		defer func() {
			fmt.Println(err)
		}()

		if err = fmt.Errorf("123"); err != nil {

		}
		return err
	}

	_ = f()
}

func TestDefer2(t *testing.T) {
	f := func() (err error) {
		defer func() {
			fmt.Println(err)
		}()

		err = fmt.Errorf("123")
		return err
	}

	_ = f()
}

func TestFunc1(t *testing.T) {
	fn := func(li ...string) {
		fmt.Println(li == nil, li)
	}

	fn()
	fn("a")
	fn("a", "b")
}
