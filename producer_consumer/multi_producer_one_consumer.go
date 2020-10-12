package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/**
场景：
多个生产者处理一批数据，然后给一个消费者消费。生产很慢，消费很快。
如果多个生产者自己生产数据，则更简单。

实现：
管道流
*/
func main() {
	var wg sync.WaitGroup
	var producerDone sync.WaitGroup

	valueCh := make(chan int, 3)
	ch := make(chan int)

	// 将数据集通过管道提供给多个生产者
	valuer := func() {
		value := []int{2, 4, 6, 8, 10}
		for _, v := range value {
			valueCh <- v
		}
		close(valueCh)
		// 生产者都结束后，关闭生产者和消费者之间的通道
		go func() {
			producerDone.Wait()
			close(ch)
		}()
	}

	producer := func() {
		for v := range valueCh {
			time.Sleep(time.Second * 1)
			ch <- v
		}
		producerDone.Done()
		wg.Done()
	}

	consumer := func() {
		for v := range ch {
			fmt.Printf("消费%s\n", strconv.Itoa(v))
		}
		wg.Done()
	}

	producerDone.Add(3)
	wg.Add(4)
	go valuer()
	go producer()
	go producer()
	go producer()
	go consumer()

	wg.Wait()
}
