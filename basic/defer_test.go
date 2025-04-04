package basic

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	x := 0
	// defer的参数是在声明时求值的，不是执行时，所以后面的x=3不会影响defer的参数
	defer fmt.Println(x)
	x = 3
}

func TestDefer2(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
