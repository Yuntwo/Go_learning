package basic

import "testing"

func TestFor(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i < len(array); i++ {
		if i&1 == 0 {
			array[i] += 1
			sum += array[i] + i
		}
	}
}
