package basic

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	x, y := 1, 2
	// 注意这里没有传递x参数，所以实际执行时的x调用的是全局变量x，实际上这里是函数闭包
	// 所以输出23
	defer func() {
		fmt.Print(x)
	}()
	x = 3
	fmt.Print(y)
}

func TestFunction1(t *testing.T) {
	x, y := 1, 2
	// 注意这里传递了x参数，所以实际执行时的x调用的是函数参数x
	// 所以输出21
	defer func(x int) {
		fmt.Print(x)
	}(x)
	x = 3
	fmt.Print(y)
}

func TestFunction2(t *testing.T) {
	x := 0
	// 这里和TestFunction1类似，都是参数传递，所以输出实际调用时x的值0
	defer fmt.Println(x)
	x = 3
}
