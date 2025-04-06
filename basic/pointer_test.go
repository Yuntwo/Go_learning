package basic

import (
	"fmt"
	"reflect"
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

func TestPointer3(t *testing.T) {
	var x *int = nil
	fmt.Println(reflect.ValueOf(x).IsNil())
}

func TestPointer4(t *testing.T) {
	var y interface{} = nil
	// 非指针类型的nil的IsNil会Panic
	fmt.Println(reflect.ValueOf(y).IsNil())
}

func TestPointer5(t *testing.T) {
	var z *interface{} = nil
	fmt.Println(reflect.ValueOf(z).IsNil())
}

func TestPointer6(t *testing.T) {
	// 直接的nil的IsNil方法会Panic
	fmt.Println(reflect.ValueOf(nil).IsNil())
}

func (a A) f4() {
	fmt.Print("f")
}

func TestPointer7(t *testing.T) {
	var a *A = nil
	// invalid memory address or nil pointer dereference
	// 无法找到指针对应的方法，无法自动解引用使用值接收者的方法，因为它本身就是nil
	a.f4()
}
