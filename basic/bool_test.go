package basic

import "testing"

func TestBool(t *testing.T) {
	// Go语言中的布尔类型只有true和false两个值
	var a = true
	if a {
		println("true")
	}
	// 这种写法是正确的，但是可以简化
	if a == true {
		println("true")
	}
}
