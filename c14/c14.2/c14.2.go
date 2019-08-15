// 通道是连接并发goroutine的管道。 您可以将值从一个goroutine发送到通道，并将这些值接收到另一个goroutine中。
// 就像一个可以用于发送类型化数据的管道，由其负责协程之间的通信
// 通道的发送/接收操作在对方准备好之前是阻塞的！！！
// 给通道赋值，必须有接收者，否则阻塞
// 拿通道值，必须有输出者，否则阻塞
// 通道的发送和接收都是原子操作：它们总是互不干扰的完成的
package c14

import (
	"fmt"
	"time"
)

// main和init其实都相当于是一个协程
func init1()  {
	// 通道的定义很简单，就是：
	// make(chan val-type)

	msg := make(chan string)

	// 使用通道 <- 语法将值发送到通道。 在这里，我们通过新的goroutine向我们上面的消息频道发送“ping”。
	go func() {
		msg <- "ping"
	}()

	// 使用 <- 语法将值将值取出

	message := <- msg

	fmt.Println(message)

	cha := make(chan int)
	go pump(cha)
	go suck(cha)
	fmt.Println(<-cha)
	time.Sleep(1e9)
}


// 通道是阻塞的，需要读取完通道值之后才能继续塞数据
func pump(ch chan int) {
	for i:= 10; ; i++ {
		ch <- i
	}
}
func suck(ch chan int)  {
	// time.Sleep(2e9)

	for {
		fmt.Println(<- ch)
	}
}

