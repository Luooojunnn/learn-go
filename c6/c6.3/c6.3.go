// 变长参数
package c7

import (
	"fmt"
)

func init()  {
	sum(1,2,3)

	// 使用切片传参
	sli1 := []int{1,2,3,4}
	sum(sli1...)

}

// 接收任意长度参数，进行相加
func sum(nums ...int){
	fmt.Println("nums",nums)
	tol := 0
	for _, v := range nums {
		tol += v
	}
	fmt.Println("tol",tol)
}