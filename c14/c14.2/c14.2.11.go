// 管道的方向
// 使用通道作为功能参数时，您可以指定通道是否仅用于发送或接收值。 这种特异性增加了程序的类型安全性。
package c14

import (
	"fmt"
)

// 通道类型可以用注解来表示它只发送或者只接收：
// var send_only chan<- int  只收
// var recv_only <-chan int  只发

func init()  {
	ping := make(chan string, 1)
	pong := make(chan string, 1)

	pings(ping, "传递值")
	pongs(ping, pong)
	fmt.Println(<- pong)
}

func pings(ping chan <- string, msg string) {
	ping <- msg
}

func pongs(ping <- chan string, pong chan <- string) {
	msg := <- ping
	pong <- msg
}