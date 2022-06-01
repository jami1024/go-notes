package copy

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	fmt.Println("使⽤range循环获取元素中的值进⾏对引用类型的深拷贝")
	slice := []int{1, 2, 3, 4}
	s2 := make([]int, 0)
	for _, v := range slice {
		s2 = append(s2, v)
	}
	s2[0] = 11
	fmt.Println(slice) //结果 [1 2 3 4]
	fmt.Println(s2)    //结果 [11 2 3 4]

	/*
		=== RUN   TestCopy
		使⽤range循环获取元素中的值进⾏对引用类型的深拷贝
		[1 2 3 4]
		[11 2 3 4]
		--- PASS: TestCopy (0.00s)
		PASS
		ok  	gonotes/copy	0.109s
	*/
}

func TestCopy2(t *testing.T) {
	fmt.Println("使⽤深拷贝数据函数: copy(⽬标切⽚,数据源)对引用类型做深拷贝")
	//copy(⽬标切⽚,数据源)深拷⻉数据函数
	s2 := []int{1, 2, 3, 4}
	s3 := []int{7, 8, 9}
	copy(s2, s3)
	s2[0] = 11
	fmt.Println(s2)
	fmt.Println(s3)

	/*
		=== RUN   TestCopy2
		使⽤深拷贝数据函数: copy(⽬标切⽚,数据源)对引用类型做深拷贝
		[11 8 9 4]
		[7 8 9]
		--- PASS: TestCopy2 (0.00s)
		PASS
		ok  	gonotes/copy	0.467s
	*/
}
