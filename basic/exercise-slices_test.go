package basic

import (
	"fmt"
	"testing"
)

// TestPic append方法会自动扩容
func TestPic(t *testing.T) {
	filteredMerchantBaseInfoList := make([]int64, 0)
	fmt.Println("Initial slice:", filteredMerchantBaseInfoList, "Len:", len(filteredMerchantBaseInfoList), "Cap:", cap(filteredMerchantBaseInfoList))

	// 切片操作 1: filteredMerchantBaseInfoList[0:0]
	slice1 := filteredMerchantBaseInfoList[0:0]
	fmt.Println("Slice [0:0]:", slice1, "Len:", len(slice1), "Cap:", cap(slice1))

	// 追加元素
	filteredMerchantBaseInfoList = append(filteredMerchantBaseInfoList, int64(1))
	fmt.Println("After append:", filteredMerchantBaseInfoList, "Len:", len(filteredMerchantBaseInfoList), "Cap:", cap(filteredMerchantBaseInfoList))
	filteredMerchantBaseInfoList = append(filteredMerchantBaseInfoList, int64(2))
	fmt.Println("After append:", filteredMerchantBaseInfoList, "Len:", len(filteredMerchantBaseInfoList), "Cap:", cap(filteredMerchantBaseInfoList))
	filteredMerchantBaseInfoList = append(filteredMerchantBaseInfoList, int64(3))
	fmt.Println("After append:", filteredMerchantBaseInfoList, "Len:", len(filteredMerchantBaseInfoList), "Cap:", cap(filteredMerchantBaseInfoList))

	// 切片操作 2: filteredMerchantBaseInfoList[3:4]
	slice2 := filteredMerchantBaseInfoList[3:4]
	fmt.Println("Slice [3:4]:", slice2, "Len:", len(slice2), "Cap:", cap(slice2))
}

func f(a []int) {
	a = append(a, 1)
}

func TestSlices(t *testing.T) {
	// len(a)是切片的长度，cap(a)是切片的容量
	// 这里创建了一个长度为0，容量为3的切片
	// 长度为0，表示切片中没有元素，a[0]无法访问，会发生runtime error
	a := make([]int, 0, 3)
	// 长度为3，容量为3多切片，有长度会有初始化0值
	//a := make([]int, 0, 3)

	for i := 0; i < 3; i++ {
		a = append(a, i)
	}
	// 这里传递的是切片的副本，函数内的修改不会影响原切片
	f(a)
	// 这里两个接收值，x打印的是value不是索引
	for _, x := range a {
		fmt.Print(x)
	}
}

func TestSlices2(t *testing.T) {
	aSlice := []A{{1}, {2}, {3}}
	ptrASlice := make([]*A, 0, len(aSlice))
	for _, a := range aSlice {
		ptrASlice = append(ptrASlice, &a)
	}

	for _, a := range ptrASlice {
		fmt.Print(a.x)
	}
}

// 等价于显式传递切片f(num []int)
func fi(num ...int) {
	num[0] = 18
}

func TestSlices3(t *testing.T) {
	i := []int{5, 6, 7}
	// 语法糖，展开为可变参数参数传递。等价于fi(i[0], i[1], i[2])
	// 注意这里传递的是原始切片的引用，函数内会直接修改原始切片的内容
	fi(i...)
	fmt.Printf("%d", i[0])
}

// slice初始化为nil，append函数可以处理nil切片，会创建一个新的切片并添加元素
func TestSlices4(t *testing.T) {
	var s []int
	s = append(s, 1)
}

// map类型初始化为nil，不能直接使用
func TestSlices5(t *testing.T) {
	var m map[string]int
	//m = make(map[string]int)
	m["one"] = 1
}

// slice初始化为nil不允许直接访问元素
func TestSlices6(t *testing.T) {
	var s []int
	fmt.Print(s[0])
}

// 可以查询nil映射的值，会返回零值（这里是0）
// 就是说对于nil映射的处理比较特殊，读操作允许默认0值，写操作不允许
func TestSlices7(t *testing.T) {
	var m map[string]int
	fmt.Print(m["one"])
}

func TestSlices8(t *testing.T) {
	x := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 7
	}(x)
	fmt.Println(x)
}

// 传递的是slice副本
func f3(a []int) {
	a = append(a, 1)
}

func TestSlices9(t *testing.T) {
	a := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		a = append(a, i)
	}
	f3(a)
	for x := range a {
		fmt.Print(x)
	}
}
