// 通道同步
package c14

import (
	"fmt"
	"time"
)

func init1()  {
	done := make(chan bool, 1)

	go worker(done)
	// 用该句进行协程阻塞，如果移除，有可能 worker 协程还没工作，就已经退出主协程了
	<- done
}

// 该方法运行在一个协程内，done 通道将被用于通知其他协程该函数运行完毕
func worker(done chan bool) {
	fmt.Println("worker 开始")
	time.Sleep(time.Second)
	fmt.Println("worker 结束")
	done <- true
}