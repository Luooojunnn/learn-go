// go 函数返回值
// go 函数支持返回多值，这个用处常被用来返回结果和错误信息
package c6

import (
	"fmt"
)

func init() {
	a, b, _ := rTwoVal()
	fmt.Println(a)
	fmt.Println(b)

	_, _, c := rTwoVal()
	fmt.Println(c)
}

func rTwoVal() (int, int, int) {
	return 1, 2, 3
}