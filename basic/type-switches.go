package basic

import "fmt"

func do(i I) {
	switch v := i.(type) {
	//case int:
	//	fmt.Printf("Twice %v is %v\n", v, v*2)
	//case string:
	//	fmt.Printf("%q is %v bytes long\n", v, len(v))
	case Ii:
		i.M()
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Ii struct {
	x int
}

func (i Ii) M() {
	fmt.Printf("Test direct implementation of interface function %x", i.x)
}

//func main() {
//	// If supporting these types, using interface{}
//	//do(21)
//	//do("hello")
//	//do(true)
//	// Otherwise need to implement M()
//	do(Ii{1})
//}
