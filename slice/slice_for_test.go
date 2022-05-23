package slice

import (
	"fmt"
	"testing"
)

// 定义切片
var names []string = []string{"a", "b", "c"}

func TestSliceForLen(t *testing.T) {
	fmt.Println("for len方式遍历切片")
	for i := 0; i < len(names); i++ {
		fmt.Printf("names的索引:%d\t值为%s\n", i, names[i])
	}
	/*
		=== RUN   TestSliceForLen
		for len方式遍历切片
		names的索引:0	值为a
		names的索引:1	值为b
		names的索引:2	值为c
		--- PASS: TestSliceForLen (0.00s)
		PASS
		ok  	gonotes/slice	0.115s
	*/
}

func TestSliceForRange(t *testing.T) {
	fmt.Println("for range方式遍历切片")

	for i := range names {
		fmt.Printf("names的索引:%d\t值为%s\n", i, names[i])
	}

	/*
		=== RUN   TestSliceForRange
		for range方式遍历切片
		names的索引:0	值为a
		names的索引:1	值为b
		names的索引:2	值为c
		--- PASS: TestSliceForRange (0.00s)
		PASS
		ok  	gonotes/slice	0.111s
	*/
}
