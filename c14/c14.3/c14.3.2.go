// channel 的范围
// 使用range函数遍历channel
package c14

import (
	"fmt"
)

func init()  {
	queue := make(chan int, 2)
	queue <- 2
	queue <- 29

	// 关闭通道
	close(queue)


	// 关闭非空管道，依然可以获取到剩余关道值
	for ele := range queue {
		fmt.Println(ele)
	}
}