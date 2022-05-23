package slice

import (
	"fmt"
	"testing"
)

func TestSliceCap(t *testing.T) {
	//cap 切片的容量，表示内存中可以存放的大小,当没有指定cap大小时，默认和长度相同
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("nums的len:%d\tnums的cap:%d\n", len(nums), cap(nums))
	nums = append(nums, 7)
	// cap满时，当增加的元素小于1024，容量就会成倍增加，若大于1024则增加0.25倍
	fmt.Printf("nums的len:%d\tnums的cap:%d\n", len(nums), cap(nums))
	nums2 := nums[1:3]
	fmt.Printf("nums2的len:%d\tnums2的值:%v\n", len(nums2), nums2)
	// 修改nums2的值nums中也会变
	nums2[0] = 100
	fmt.Printf("nums2的值:%v\tnums的值:%v\n", nums2, nums)
	/*
		=== RUN   TestSliceCap
		nums的len:6	nums的cap:6
		nums的len:7	nums的cap:12
		nums2的len:2	nums2的值:[2 3]
		nums2的值:[100 3]	nums的值:[1 100 3 4 5 6 7]
		--- PASS: TestSliceCap (0.00s)
		PASS
		ok  	gonotes/slice	0.241s
	*/
}
