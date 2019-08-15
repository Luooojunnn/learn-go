// 闭包
package c6

import (
	"fmt"
)

func init()  {
	next := intSeq()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	next2 := intSeq()

	fmt.Println(next2())
	fmt.Println(next2())
	fmt.Println(next2())

	fmt.Println(a())

	
}

func intSeq() func() int  {
	i := 0
	return func() int {
		i++
		return i
	}
}

func a() (r int) {
	defer func() int {
		return 100
	}()
	return 1
}

