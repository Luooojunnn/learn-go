// 计时器（Ticker）

package c14

import (
	"fmt"
	"time"
)

func init3()  {
	ticker := time.NewTicker(500 * time.Millisecond)

	go func(){
		for i := range ticker.C {
			fmt.Println(i)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("ticker停止")
}
