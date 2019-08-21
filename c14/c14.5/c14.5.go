// 超时

package c14

import (
	"fmt"
	"time"
)

func init1()  {
	c1 := make(chan string, 1)
	// 模拟 io 操作，两秒延时
	go func()  {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// 此处做一个竞态，1s没结果就走1s的结果
	select {
	case msg := <- c1:
		fmt.Println(msg)
	case <- time.After(1 * time.Second):
		fmt.Println("1s超时了，老弟")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case msg := <- c2:
		fmt.Println(msg)
	case <- time.After(3 * time.Second):
		fmt.Println("3s超时了，老弟")
	}
}