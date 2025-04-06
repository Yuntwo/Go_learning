package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestChan_Closed(t *testing.T) {
	// 无缓冲通道
	ch := make(chan string)
	ch <- "1"
	// 主goroutine立即向channel发送"1"后，但由于没有接收者（消费者goroutine还未启动），发送"2"操作会阻塞
	ch <- "2"
	// 由于主goroutine被阻塞，消费者goroutine永远不会被执行
	go func() {
		for m := range ch {
			fmt.Print(m)
		}
	}()
	// 最终程序会因为所有goroutine都阻塞而报死锁错误(panic)
}

func TestChan(t *testing.T) {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i + 1
		}
		// 关闭通道不是必须的，仍然可以从关闭的通道中读取数据，通道为空读出为0值
		close(ch)
	}()
	// 这种情况下关闭通道只是为了通知range没有更多数据，可以退出循环了
	// 仍然可以从关闭的通道中读取数据
	for m := range ch {
		fmt.Print(m)
	}
}

// Panic: send on closed channel
func TestChan2(t *testing.T) {
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- idx
		}(i)
	}

	fmt.Print(<-ch)
	close(ch)
	time.Sleep(2 * time.Second)
}

func TestChan3(t *testing.T) {
	// 无缓冲通道必须有接收者才能发送数据，否则会阻塞
	ch := make(chan interface{})
	x := 1
	// 程序会阻塞在select处
	// 死锁Panic，不会走到任何一个
	select {
	case ch <- x:
		println("100")
	case ch <- x:
		println("200")
		return
	}
}

func TestChan4(t *testing.T) {
	ch := make(chan int) // 无缓冲channel
	// 启动3个goroutine，不会阻塞主goroutine
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- idx // 每个goroutine尝试发送自己的索引到channel
		}(i)
	}
	fmt.Print(<-ch) // 从channel接收一个值并打印
	// panic: send on closed channel
	close(ch)
	time.Sleep(2 * time.Second)
}
