package point

import (
	"fmt"
	"testing"
)

func TestArrayPoint(t *testing.T) {
	fmt.Println("指针数组")

	//定义四个变量
	a, b, c, d := 1, 2, 3, 4
	arr1 := [4]int{a, b, c, d}
	arr2 := [4]*int{&a, &b, &c, &d} //将所有变量的指针，放进arr2⾥⾯

	fmt.Printf("arr1的值%v\n", arr1) //arr1的值[1 2 3 4]

	fmt.Printf("数组指针arr2%v\n", arr2) // 数组指针arr2[0xc0001161e0 0xc0001161e8 0xc0001161f0 0xc0001161f8]

	arr1[0] = 100 //修改arr1中的值

	fmt.Println("修改后arr1的值：", arr1) // [100 2 3 4]

	fmt.Println("a=", a) //变量a的值还是1，相当于值传递，只修改了数值的副本。

	//修改指针数组
	*arr2[0] = 100 //修改指针的值
	fmt.Printf("修改后的arr2%v\n", arr2)
	fmt.Println("a=", a) //100

	//循环数组，⽤*取数组中的所有值。
	//引⽤传递 修改的是内存地址所对应的值 所以a也修改了

	for i := 0; i < len(arr2); i++ {
		fmt.Println(*arr2[i])
	}
}
