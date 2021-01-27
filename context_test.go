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
