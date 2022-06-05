package point

import (
	"fmt"
	"testing"
)

func TestPointArray(t *testing.T) {
	fmt.Println("数组指针")

	//创建⼀个普通的数组
	// var arr [3]int = [3]int{1, 2, 3}
	arr := [4]int{1: 1, 2: 2, 3: 3}
	fmt.Printf("数组%v\n", arr)

	//创建⼀个指针 ⽤来存储数组的地址即：数组指针
	var p *[4]int             //将数组arr的地址，存储到数组指针p上。
	fmt.Printf("数组指针%v\n", p) //数组的指针 &[1 2 3] 后⾯跟数组的内容
	p = &arr
	//获取数组指针中的具体数据 和数组指针⾃⼰的地址
	fmt.Printf("指针p的值%v\n", *p)  //指针所对应的数组的值
	fmt.Printf("指针p的地址%v\n", &p) //该指针⾃⼰的地址0xc00000e048

	//修改数组指针中的数据
	(*p)[0] = 100
	fmt.Println(arr) //[100 1 2 3]

	//简化写法
	p[1] = 200
	fmt.Println(arr) //结果： [100 200 2 3]
}
