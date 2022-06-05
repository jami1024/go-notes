package fuc_test

import (
	"fmt"
	"testing"
)

func TestFucStatement(t *testing.T) {
	fmt.Println("函数的声明")
	func1() // 执行

}

//func关键字加上函数名、参数列表、返回值
func func1() {
	fmt.Println("函数")
}
