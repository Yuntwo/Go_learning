# Note
## Grammar
- 读取数据用`<-ch`，规范<-和ch之间不需要空格，加了空格也是合法的
- 发送数据用`ch <- 1`，一般有空格
- 读取数据到未声明变量以及是否有更多值用`v, ok := <-ch`
## Buffered Channels
- The default channel is unbuffered, that is like make(chan int, 0). This means you can't send and read it in a single routine that only introduces deadlock.
## Send to and receive from channel
- Avoid deadlock when receiving from an empty channel and sending to a full channel, which means you must ensure the equal number of send and receive operations.
## Close a channel
- Only the sender should close a channel, never the receiver.
  - Sending on a closed channel will cause a panic.
  - Receiver can still read from a closed channel, and read from a closed empty channel won't be blocked but receives (0, false).
-  Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.
## Select multiple channels (应该只适用于channel语法)，主要用于多路复用
- Default通常用于非阻塞操作
- select{}会永久阻塞，通常用于阻止程序退出
## chan是引用类型，参数传递时不用指针，直接传

# Mutex(通常是资源)
- 互斥锁是一种常见的控制并发访问的方法，可以用来保证在某一时刻只有一个goroutine访问共享资源
- 一个需要并发访问的类型通常会有一个互斥锁作为字段
```
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

var lock = sync.Mutex{}
```

