package basic

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	// 在select语句中，每次循环都会重新创建两个新的计时器，但只使用一次，会造成资源泄漏
	// 每次调用time.After都会创建一个新的计时器通道
	for {
		select {
		// 未触发的计时器会留在内存中直到超时，会造成资源泄漏
		case <-time.After(time.Second * 10):
			println("timer 10 secs")
		// 每次都是1秒的计时器先触发
		case <-time.After(time.Second * 1):
			println("timer 1 secs")
		}
	}
}

// 正确触发写法
func TestTimer2(t *testing.T) {
	timer10 := time.After(time.Second * 10)
	timer1 := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-timer10:
			println("timer 10 secs")
			timer10 = time.After(time.Second * 10) // 重置10秒计时器
		case <-timer1.C:
			println("timer 1 secs")
		}
	}
}
