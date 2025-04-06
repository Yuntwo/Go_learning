package basic

import (
	"encoding/json"
	"fmt"
	"testing"
)

type MyData struct {
	One int `json:"one"` // 公开字段(大写开头)，有json标签
	two int `json:"two"` // 私有字段(小写开头)，有json标签
}

func TestJson(t *testing.T) {
	var out MyData
	_ = json.Unmarshal([]byte(`{"one": 1, "two": 2}`), &out)
	fmt.Printf("%+v\n", out)
}
