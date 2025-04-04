package basic

import (
	"fmt"
	"testing"
	"unsafe"
)

type People struct {
	Name string
	Age  int64
}

// 值接收者方法
func (p People) SetAge(age int64) {
	p.Age = age // 修改的是副本
}

// 指针接收者方法
func (p *People) PSetAge(age int64) {
	// 自动解引用指针(*p)修改原对象
	// 修改的是原对象
	p.Age = age
}

func TestPointer(t *testing.T) {
	p := &People{}
	// 使用值接收者，Go会自动解引用指针(*p)创建副本
	// 修改的是副本的Age，原对象的Age不变
	p.SetAge(1)
	// 这里自动解引用指针(*p)获取原对象的值
	fmt.Println(p.Age)
	// 使用指针接收者，直接操作原对象
	p.PSetAge(1)
	// 这里自动解引用指针(*p)获取修改后原对象的值
	fmt.Println(p.Age)
}

type Demo struct {
	Age    int
	Name   string
	Gender bool
}

// 当方法不访问接收者的字段时，即使接收者为nil也不会panic
func (*Demo) Print() {
	fmt.Print("test")
}

func TestPointer2(t *testing.T) {
	d := (*Demo)(unsafe.Pointer(nil))
	// 虽然d是nil，但Go允许在nil指针上调用方法
	// Go的方法调用实际上是函数调用，接收者作为第一个参数传递
	d.Print()
}
