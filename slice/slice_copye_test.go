package slice

import (
	"fmt"
	"testing"
)

func TestSliceCopy(t *testing.T) {
	fmt.Println("切片的copy")
	n1 := []int{0, 1, 2, 3, 4, 5}
	n2 := []int{10, 11, 12, 13, 14, 15, 16}

	//copy(dst, src) src => dst,如果源比目标多，那么多出来的值并不会copy到目标中

	copy(n1, n2)
	fmt.Printf("n1 copy n2 后的结果%v\n", n1)

	/*
		=== RUN   TestSliceCopy
		切片的copy
		n1 copy n2 后的结果[10 11 12 13 14 15]
		--- PASS: TestSliceCopy (0.00s)
		PASS
		ok  	gonotes/slice	(cached)

	*/
}
