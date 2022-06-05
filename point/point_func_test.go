package point

import (
	"fmt"
	"testing"
)

func TestPointFunc(t *testing.T) {
	fmt.Println("指针函数")
	a1 := fun1()
	fmt.Printf("a1的类型：%T,a1的地址是%p 数值为%v \n", a1, &a1, a1) //[]int,a1的地址是0xc00000c048 数值为[1 2 3]

	a2 := fun2()
	fmt.Printf("a2的类型：%T,a2的地址是%p 数值为%v \n", a2, &a2, a2) //*[]int,a2的地址是0xc00000e048 数值为&[1 2 3 4]
	fmt.Printf("a2的值为：%p\n", a2)                          //0xc00000c090 指针函数返回的就是指针
}

//⼀般函数
func fun1() []int {
	c := []int{1, 2, 3}
	return c
}

//指针函数 返回指针
func fun2() *[]int {
	c := []int{1, 2, 3, 4}
	fmt.Printf("c的地址为%p：\n", &c)
	return &c

}
