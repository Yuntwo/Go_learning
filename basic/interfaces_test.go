package basic

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestAbser_Abs(t *testing.T) {
	var a Abser
	// Use () to instantiate a primitive type based type
	f := MyFloat(-math.Sqrt2)
	// f has primitive type float64
	fmt.Println(f)
	// Use {} to instantiate a struct type based type
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// Error
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())
}

type Binary uint64

// 没有方法的类型使用eface的底层数据结构存储
func TestEface(t *testing.T) {
	b := Binary(200)
	// 这个类型转换使用了Go运行时（runtime包）将具体类型转换为interface{}的convT2E函数
	// 核心逻辑是：
	//func convT2E(t *_type, elem unsafe.Pointer) (e eface) {
	//	e._type = t
	//	e.data = elem
	//	return
	//}
	// 这个过程保证了动态类型信息和数据能够正确存储到interface{}中
	any := (interface{})(b)
	fmt.Println(any)                 // 输出200
	fmt.Println(reflect.TypeOf(b))   // basic.Binary
	fmt.Println(reflect.TypeOf(any)) // basic.Binary
}
