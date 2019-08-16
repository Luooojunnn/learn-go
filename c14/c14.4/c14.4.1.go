// 非阻塞频道操作
package c14

import (
	"fmt"
)

// 通道上的基本发送和接收是阻塞的。 但是，我们可以使用带有default子句的select来实现非阻塞发送，接收甚至非阻塞的多路选择。

func init()  {
	msg := make(chan string)
	bol := make(chan bool)

	// 无缓冲channel被阻塞了，所以都走的default

	select {
	case m := <- msg:
		fmt.Println("收到msg:", m)
	default:
		fmt.Println("没有收到msg")
	}

	m := "nihao"
	select {
	case msg <- m:
		fmt.Println("收到msg:", m)
	default:
		fmt.Println("没有收到msg")
	}

	select {
	case m := <- msg:
		fmt.Println("收到msg:", m)
	case b := <- bol:
		fmt.Println("收到bool:", b)
	default:
		fmt.Println("没有收到msg&bool")
	}
}