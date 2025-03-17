package basic

import (
	"math"
)

// Abser 定义接口
type Abser interface {
	Abs() float64
}

// MyFloat 实现接口1，float64自定义类型
// 实现接口需要通过类型隐式实现，任何类型都可实现接口
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex 实现接口2，接口体自定义类型
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Error
// interface type cannot be receiver, that is cannot implement the method of type interface directly
//func (f Abser) Abs() float64 {
//  if f < 0 {
//      return float64(-f)
//  }
//  return float64(f)
//}
