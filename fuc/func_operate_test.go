package fuc_test

import (
	"fmt"
	"testing"
)

func TestFuncOperate(t *testing.T) {
	num := getSum() //函数的调⽤
	fmt.Printf("100以内之和为%d\n", num)
}

//定义函数 求1-100的和
func getSum() int {
	sum := 0
	for i := 0; i <= 50; i++ {
		sum += i
	}
	return sum
}
