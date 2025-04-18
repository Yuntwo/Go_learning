package basic

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// Return type no literal
func run() error {
	// Only &MyError type implemented Error() so that can be called using Println
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

//func main() {
//	if err := run(); err != nil {
//		fmt.Println(err)
//	}
//}
