package basic

import (
	"fmt"
	"math"
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
