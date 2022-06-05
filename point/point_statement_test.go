package point

import (
	"fmt"
	"testing"
)

// int float string 数组 都是属于值类型

func TestPointStatement(t *testing.T) {
	fmt.Println("指针的定义")
	//定义⼀个变量
	var a int = 2
	fmt.Printf("变量a的地址为%p\n", &a) //通过%p占位符, &符号获取变量的内存地址。变量A的地址为0xc0000162d0

	//创建⼀个指针
	//指针的声明 通过 *T 表示T类型的指针
	var i *int                  //int类型的指针
	var f *float64              //float64类型的指针
	fmt.Printf("指针i得值:%v\n", i) // < nil >空指针
	fmt.Printf("指针f得值:%v\n", f) // < nil >空指针

	//因为指针存储的变量的地址 所以指针存储值
	i = &a
	fmt.Printf("i赋值后为:%v\n", i) //i存储a的内存地址0xc0000162d0
	fmt.Println(*i)             //i存储这个指针存储的变量的数值2
	*i = 100
	fmt.Println(*i) //100
	fmt.Println(a)  //100通过指针操作 直接操作的是指针所对应的数值
}
