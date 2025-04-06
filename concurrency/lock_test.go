package concurrency

import (
	"sync"
	"testing"
)

type myMutex sync.Mutex

func TestLock1(t *testing.T) {
	// 编译错误，因为myMutex没有实现Lock和Unlock方法
	// myMutex是新定义的类型，虽然底层类型是sync.Mutex，但它不会继承sync.Mutex的方法
	//var mtx myMutex
	//mtx.Lock()
	//mtx.Unlock()
}

type myLocker struct {
	// 这里使用了结构体嵌入sync.Mutex，通过嵌入方式，myLocker会继承sync.Mutex的所有方法
	sync.Mutex
}

func TestLock2(t *testing.T) {
	// 默认初始化为0，表示没有锁定，而不是指针类型的nil
	var lock myLocker
	lock.Lock()
	lock.Unlock()
}

// sync.Locker是一个接口，它只有Lock和Unlock方法，所以myLocker1可以实现sync.Locker接口
type myLocker1 sync.Locker

func TestLock3(t *testing.T) {
	// sync.Mutex实现了sync.Locker接口，所以可以声明为myLocker1
	// 实际对象是sync.Mutex，实现了Lock和Unlock方法
	var lock myLocker1 = new(sync.Mutex)
	lock.Lock()
	lock.Unlock()
}
