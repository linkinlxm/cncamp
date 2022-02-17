//队列：
//队列长度 10，队列元素类型为 int
//生产者：
//每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
//消费者：
//每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	//defer close(ch1)
	//for {
	go producer(ch1)
	// sleep 15秒 观察阻塞现象
	time.Sleep(15 * time.Second)
	go consumer(ch1)
	time.Sleep(100 * time.Second)
	//}
}

//producer
func producer(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 50; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		fmt.Printf("Produce data with value as : %d\n", i)
	}
}

// consumer
func consumer(ch <-chan int) {
	time := time.NewTimer(5 * time.Second)
	select {
	case <-ch:
		for value := range ch {
			fmt.Printf("Receive from channel wit msg: %d\n", value)
		}
	case <-time.C:
		//可以在main函数只保留go consumer(ch1)，然后后面加个延时15s就可以验证这里的timeout报错
		fmt.Printf("[WSRN] Consumer timeout waiting from channel ch")
	}
	//for value := range ch {
	//	fmt.Printf("Receive from channel wit msg: %d\n", value)
	//}
}
