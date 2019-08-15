// 递归函数
package c6

import (
	"fmt"
)

func init() {
	num := 10
	fmt.Printf("斐波那契数列 %d 的值是 %d", num, fib(num))
}

func fib(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	
	return fib(num-1) + num
}