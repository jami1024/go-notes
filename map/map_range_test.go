package map_test

import (
	"fmt"
	"testing"
)

func TestMapRange(t *testing.T) {
	fmt.Println("map的遍历")
	map1 := make(map[int]string)

	map1[1] = "a"

	map1[2] = "b"

	map1[3] = "c"

	map1[4] = "d"

	//遍历map
	for key, val := range map1 {
		fmt.Println(key, val)
	}
	/*
		map的遍历
		3 c
		4 d
		1 a
		2 b
	*/
}
