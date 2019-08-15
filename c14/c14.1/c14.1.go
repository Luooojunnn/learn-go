// Go 语言为构建并发程序的基本代码块是 协程 (goroutine) 与通道 (channel)
// 不要通过共享内存来通信，而通过通信来共享内存。
// 可以理解为 async 关键字，使任务异步化，只是是使用协程进行的异步
// 用来启动这个协程的函数不会得到任何的返回值。
package c14

import (
	"fmt"
)

func init()  {
	// 常用方式，同步调用
	f("direct")

	// 协程的使用方式：go
	go f("goroutine")

	// 也可以将匿名函数放在协程中
	go func(msg string) {
		fmt.Println(msg)
	}("nihao")

	// 以上两个方法运行在两个分离的协程当中
	fmt.Scanln()
	fmt.Println("done!")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}