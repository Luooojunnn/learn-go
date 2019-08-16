// 关闭通道
// 关闭通道对于管道的接收方来说是很有用的
package c14

import (
	"fmt"
)

func init()  {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 异步的，输出顺序完全不固定

	go func() {
		for {
			v, ok := <- jobs
			if ok {
				fmt.Println("收到数据",v)
			}else {
				fmt.Println("数据全部接收完毕")
				done <- true
				return 
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Printf("发送了数据%d\n", i)
	}
	close(jobs)
	fmt.Println("发送完毕数据")

	
	// 使用该方法阻塞主线程，进行同步等待
	<- done

}