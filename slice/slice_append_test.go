package slice

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	// append新增
	// 定义切片,注意:=这种方式不能定义全局变量
	names := []string{"a", "b", "c"}

	fmt.Printf("names:%q\n", names)
	names = append(names, "d")
	fmt.Printf("names:%q\n", names)

	/*
		=== RUN   TestSliceAppend
		names:["a" "b" "c"]
		names:["a" "b" "c" "d"]
		--- PASS: TestSliceAppend (0.00s)
		PASS
		ok  	gonotes/slice	0.296s
	*/
}


