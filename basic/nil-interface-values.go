package basic

import "fmt"

type I interface {
	M()
}

//func main() {
//	var i I
//	describe(i)
//	// If an interface has no implementations of its members, then called nil interface
//	// This will introduce error when calling its member
//	i.M()
//}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
