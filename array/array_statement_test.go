package array

import (
	"fmt"
	"testing"
)

func TestArrayStatement(t *testing.T) {
	fmt.Println("数组的声明")
	var names [3]string = [3]string{"zhangsan1", "zhangsan2"}
	var nums [5]int
	fmt.Printf("names类型%T,\tnums类型%T\n", names, nums)
	fmt.Println(names, nums)
	/*
		=== RUN   TestArrayStatement
		数组的声明
		names类型[3]string,	nums类型[5]int
		[zhangsan1 zhangsan2 ] [0 0 0 0 0]
		--- PASS: TestArrayStatement (0.00s)
		PASS
		ok  	gonotes/array	0.485s
	*/
}
