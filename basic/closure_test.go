package basic

import (
	"fmt"
	"testing"
)

type Order struct {
	Status int
}

func updateOrder(order *Order) {
	// 注意传递都是值传递，只是值传递的是指针，也就是地址值
	// 直接修改这个地址变量的值是无效的
	txUpdateFc := func() {
		newOrder := &Order{Status: 2}
		order = newOrder // 重新赋值 order 变量
	}
	txUpdateFc()
	fmt.Println("Inside updateOrder:", order.Status)
}

func updateOrderVal(order *Order) {
	// 注意传递都是值传递，只是值传递的是指针，也就是地址值
	// 直接通过地址修改值是有效的
	txUpdateFc := func() {
		order.Status = 2 // 重新赋值 order 变量
	}
	txUpdateFc()
	fmt.Println("Inside updateOrder:", order.Status)
}

func TestClosure(t *testing.T) {
	originalOrder := &Order{Status: 1}
	updateOrder(originalOrder)
	fmt.Println("In main:", originalOrder.Status)
}

func TestClosureVal(t *testing.T) {
	originalOrder := &Order{Status: 1}
	updateOrderVal(originalOrder)
	fmt.Println("In main:", originalOrder.Status)
}
