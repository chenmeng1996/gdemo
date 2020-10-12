package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int, 3)

	producer := func() {
		value := []int{2, 4, 6}
		for _, v := range value {
			time.Sleep(time.Second * 1)
			ch <- v
		}
		close(ch)
		wg.Done()
	}

	consumer := func() {
		for v := range ch {
			fmt.Printf("消费%s\n", strconv.Itoa(v))
		}
		wg.Done()
	}

	wg.Add(2)
	go producer()
	go consumer()

	wg.Wait()
}
