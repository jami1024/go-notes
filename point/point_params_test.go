package point

import (
	"fmt"
	"testing"
)

func TestPointParams(t *testing.T) {
	fmt.Println("指针参数")
	s := 10
	fmt.Println(s) //调⽤函数之前数值是10
	// fun1(&s)
	funcPointInt(&s)
	fmt.Println(s) //调⽤函数之后再访问则被修改成2
}

//接收⼀个int类型的指针作为参数
func funcPointInt(a *int) {
	*a = 2
}
