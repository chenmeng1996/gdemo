package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	handle := func(ctx context.Context, duration time.Duration) {
		select {
		case <-ctx.Done():
			fmt.Println("handle", ctx.Err())
		case <-time.After(duration):
			fmt.Println("process request with", duration)
		}
	}

	// 模拟1秒的逻辑处理
	go handle(ctx, 1*time.Second)

	select {
	case <-ctx.Done():
		// 其他goruntine超时中断
		fmt.Println("main", ctx.Err())
	}
}

// 中断1个子goroutine
func TestContest2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		defer func() {
			fmt.Println("子goroutine结束")
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("子goroutine被中断")
				return
			default:
				fmt.Println("子goroutine正在执行")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(10 * time.Second)
}

// 中断多个平级子goroutine
func TestContest3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	handler := func(ctx context.Context, name string) {
		defer func() {
			fmt.Println("子goroutine结束")
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("子goroutine被中断")
				return
			default:
				fmt.Println(name + ": 子goroutine正在执行")
				time.Sleep(2 * time.Second)
			}
		}

	}

	go handler(ctx, "1")
	go handler(ctx, "2")
	go handler(ctx, "3")

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(10 * time.Second)
}

// 中断多个嵌套子goroutine
func TestContest4(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	handler := func(ctx context.Context, name string) {
		defer func() {
			fmt.Println("子goroutine结束")
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("子goroutine被中断")
				return
			default:
				fmt.Println(name + ": 子goroutine正在执行")
				time.Sleep(2 * time.Second)
			}
		}
	}

	go func() {
		go handler(ctx, "2")
		handler(ctx, "1")
	}()

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(10 * time.Second)
}
