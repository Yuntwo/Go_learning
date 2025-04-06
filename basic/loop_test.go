package basic

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i < len(array); i++ {
		if i&1 == 0 {
			array[i] += 1
			sum += array[i] + i
		}
	}
}

func TestLoop(t *testing.T) {
	x := []string{"a", "b", "c"}
	// range循环在遍历切片时一个接收值返回的是索引(index)，而不是元素值(value)
	for v := range x {
		// 遍历元素值应该使用两个变量接收range的返回值
		//for i, v := range x {  // i 是索引，v 是元素值
		fmt.Print(v)
	}
}
