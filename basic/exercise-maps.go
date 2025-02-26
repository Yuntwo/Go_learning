package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	ss := strings.Fields(s)
	m := make(map[string]int)
	for _, sss := range ss {
		value, has := m[sss]
		if has {
			m[sss] = value + 1
		} else {
			m[sss] = 1
		}
	}
	return m
}

//func main() {
//	wc.Test(WordCount)
//}
