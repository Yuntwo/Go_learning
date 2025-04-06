package basic

import (
	"fmt"
	"testing"
)

type A struct {
	x int
}

func TestStruct3(t *testing.T) {
	// map里的值类型是不可寻址的，返回的是一个副本，不能直接修改
	m := map[string]A{
		"x": {1},
	}
	// 编译报错
	//m["x"].x = 2
	// 正确写法
	tmp := m["x"]
	tmp.x = 2
	m["x"] = tmp
}

func TestStruct4(t *testing.T) {
	// 这种都是缺省初始化struct类型A的写法
	s := []struct{ x int }{{1}, {2}, {3}}
	s[0].x = 1
}

func TestStruct5(t *testing.T) {
	// 这种都是缺省初始化struct类型A的写法
	m := map[string]*A{
		"x": {1},
	}
	// 指针类型可以直接修改
	m["x"].x = 2
}

func TestStruct6(t *testing.T) {
	// 这里初始化了一个空的map，就是m不是nil，但是map的value是nil，所以m["x"]是nil，不能直接访问x字段
	m := map[string]*A{}
	m["x"].x = 2
}

func (a A) f1() {
	a.x = 2
}

func TestStruct7(t *testing.T) {
	a := &A{x: 1}
	// 不会报错，只是实际修改不生效
	a.f1()
	fmt.Print(a.x)
}

func (a *A) f6() {
	fmt.Print("f6")
}

func TestStruct8(t *testing.T) {
	// 空指针方法只要不访问字段就不会报错
	var a *A = nil
	a.f6()
}

func TestStruct9(t *testing.T) {
	var a A
	// 会自动引用为指针类型的方法，因为这里a是值类型会初始化为零值
	a.f6()
}
