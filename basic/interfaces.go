package main

import (
	"math"
)

type Abser interface {
	Abs() float64
}

//func main() {
//    var a Abser
//    // Use () to instantiate a primitive type based type
//    f := MyFloat(-math.Sqrt2)
//    // f has primitive type float64
//    fmt.Println(f)
//    // Use {} to instantiate a struct type based type
//    v := Vertex{3, 4}
//
//    a = f  // a MyFloat implements Abser
//    a = &v // a *Vertex implements Abser
//
//    // In the following line, v is a Vertex (not *Vertex)
//    // and does NOT implement Abser.
//    //a = v
//
//    fmt.Println(a.Abs())
//}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// interface type cannot be receiver, that is cannot implement the method of type interface directly
//func (f Abser) Abs() float64 {
//   if f < 0 {
//       return float64(-f)
//   }
//   return float64(f)
//}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
