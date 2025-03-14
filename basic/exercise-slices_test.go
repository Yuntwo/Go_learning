package main

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
