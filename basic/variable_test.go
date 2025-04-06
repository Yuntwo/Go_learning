package basic

import (
	"fmt"
	"testing"
)

// x被声明为变量，类型推断为int
var x = 1

func TestVariable(t *testing.T) {
	x := 2
	fmt.Print(x)
}

func TestVariable2(t *testing.T) {
	// 不能直接将int类型赋值给int64和float64变量，需要显式类型转换
	//var y int64 = x
	//var z float64 = x
	//fmt.Println(y, z)
}

func TestVariable3(t *testing.T) {
	// x是未类型化常量(untyped constant)，值可以根据上下文自动转换为需要的类型，这是Go语言中常量的一个有用特性
	const x = 2
	// 可以直接赋值给int64和float64变量
	var y int64 = x
	var z float64 = x
	fmt.Println(y, z)
}

func TestVariable4(t *testing.T) {
	// x是变量，可以获取其内存地址
	var x = 2
	fmt.Println(&x)
}

func TestVariable5(t *testing.T) {
	// 编译错误，x是常量，常量没有内存地址，不能对常量使用&取地址操作符
	//const x = 2
	//fmt.Println(&x)
}

func TestVariable6(t *testing.T) {
	x := 1
	{
		x := 2

		fmt.Print(x)
	}

	fmt.Print(x)
}

func TestVariable7(t *testing.T) {
	// 这里函数内相当于{}局部作用域，输出2
	x := 2
	fmt.Print(x)
}
