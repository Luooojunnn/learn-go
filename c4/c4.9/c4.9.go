// 指针
// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
// 一个指针变量可以指向任何一个值的内存地址
package c4

import (
	"fmt"
)

func init()  {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("副本函数zeroval过后:", i)

	zeroptr(&i)
	fmt.Println("指针函数zeroptr过后:", i)

	fmt.Println("i的指针:", &i)
	
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(ival *int) {
	*ival = 0
}