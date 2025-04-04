package basic

import (
	"strconv"
	"strings"
	"testing"
)

func TestMath(t *testing.T) {
	money := 566000
	println(strings.TrimRight(strings.TrimRight(strconv.FormatFloat(float64(money)/100.0, 'f', 2, 64), "0"), ".") + "元")
}
