package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// 当主goroutine退出时，整个程序结束，所有未完成的goroutine（包括子goroutine）都会被强制终止
func TestRoutine(t *testing.T) {

	go func() {
		fmt.Print("go")
		time.Sleep(3 * time.Second)
		fmt.Print("go done")
	}()
	time.Sleep(1 * time.Second)
	fmt.Print("done")
	// 可以用sync.WaitGroup等待所有goroutine完成
}
