// 带缓冲的通道
// 默认无缓冲通道是只能存储一个值的
// 典型死锁情形
// 1.  创建了一个无缓冲的channel

// 2.  主routine要向channel中放入一个数据，但是因为channel没有缓冲，相当于channel一直都是满的，所以这里会发生阻塞。可是下面的那个goroutine还没有创建呢，主routine在这里一阻塞，整个程序就只能这么一直阻塞下去了，然后。。。然后就没有然后了。。死锁！

// ※从这里可以看出，对于无缓冲的channel，放入操作和取出操作不能再同一个routine中，而且应该是先确保有某个routine对它执行取出操作，然后才能在另一个routine中执行放入操作。


package c14

import (
	"fmt"
	// "time"
)

func init2()  {
	// 缓冲数为2的int型通道
	// msg := make(chan int, 2)

	// msg <- 1
	// msg <- 2

	// fmt.Println(<- msg)
	// fmt.Println(<- msg)
	
	// 演示阻塞
	in := make(chan int)
	// 对于无缓冲通道，千万不要把取和放 都在同一个 goroutine 里边，肯定会死锁
	// go func(in chan int) {
	// 	fmt.Println(<-in)
	// }(in)
	go func(in chan int) {
		in <- 100
	}(in)
	
	fmt.Println(<-in)
	
	// time.Sleep(1e9)

}

// func a(in chan int) {
// 	fmt.Print( <- in)
// }