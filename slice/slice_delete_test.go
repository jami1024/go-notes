package slice

import (
	"fmt"
	"testing"
)

func TestSliceDelete(t *testing.T) {
	fmt.Println("切片的删除")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	fmt.Printf("nums的长度为:%d\t元素为%v\n", len(nums), nums)
	//删除索引为0的值
	fmt.Println(nums[1:len(nums)])

	/*
		=== RUN   TestSliceDelete
		切片的删除
		nums的长度为:16	元素为[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
		[2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
		--- PASS: TestSliceDelete (0.00s)
		PASS
		ok  	gonotes/slice	0.108s


	*/
}
