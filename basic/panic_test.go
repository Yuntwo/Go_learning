package basic

import (
	"fmt"
	"sync"
	"testing"
)

func TestPanic(t *testing.T) {
	// 无效
	recover()
	// panic("not good") 触发 panic
	panic("not good")
	// 无效
	recover()
	// 程序退出不可达
	fmt.Print("ok")
}

func TestPanic2(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			println("oh no")
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		panic("panic")
		// 注意因为这里panic之后wg.Done()不会被执行，主goroutine的wg.Wait()会被阻塞，defer中的recover()不会被执行
		wg.Done()
	}()

	wg.Wait()
	println("done")
}
