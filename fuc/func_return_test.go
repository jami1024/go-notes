package fuc_test

import (
	"fmt"
	"testing"
)

func TestFucnReturn(t *testing.T) {
	fmt.Println("函数的返回值")
	num, ok := getnum()
	fmt.Println(num, ok)
}

// 返回值类型数字和字符串
func getnum() (int, string) {
	return 12, "ok"
}
