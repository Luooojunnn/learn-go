// 函数
// go 语言的核心
package c6

import (
	"fmt"
)

type multType func(int,int) int

func init()  {
	// 两数相加
	sum(1,2)

}

// 函数被调用的时候，这些实参将被复制（简单而言）然后传递给被调用函数

func sum(a,b int) float32 {
	fmt.Println(float32(a + b))
	return float32(a + b)
}