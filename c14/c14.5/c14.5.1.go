// Timers

package c14

import (
	"fmt"
	"time"
)

func init2()  {
	time1 := time.NewTimer(2 * time.Second)

	<- time1.C

	fmt.Println("time1到")

	time2 := time.NewTimer(time.Second)

	go func()  {
		<- time2.C
		fmt.Println("time2到")
	}()

	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("stop2到")
	}

}