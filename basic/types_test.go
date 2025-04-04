package basic

import (
	"fmt"
	"reflect"
	"testing"
)

type F interface {
	f()
}

type A struct {
	x int
}

func (a A) f() {
}

func TestTypes(t *testing.T) {
	var a interface{} = A{}
	fmt.Println(reflect.TypeOf(a))
	// a.(type)只能用于type switch，用于检查interface{}变量的实际类型
	switch a.(type) {
	// type switch只会执行第一个匹配的case，这里a实际上既是A类型也是F类型，哪个case放在前面输出哪一个
	case F:
		fmt.Print("F")
	case A:
		fmt.Print("A")
	default:
		fmt.Print("default")
	}
}
