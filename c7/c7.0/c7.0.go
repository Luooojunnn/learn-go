// 数组相关
// 数组在 go 中是值类型

package c7

import (
	"fmt"
)

func init(){
	var arr1 [5]int
	fmt.Println(arr1)

	arr1[4] = 100
	fmt.Println("arr1", arr1)
	fmt.Println("arr1[4]", arr1[4])
	fmt.Println("arr1的长度", len(arr1))


	arr2 := [5]int{1,2,3,4,5}
	fmt.Println("arr2", arr2)
	
	// 二维数组
	// 打印一个3*3星阵
	var arr3 [3][3]int
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}