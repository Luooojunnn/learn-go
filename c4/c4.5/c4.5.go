// 4.5 基本类型和运算符

// 在格式化字符串里，%d 用于格式化整数（%x 和 %X 用于格式化 16 进制表示的数字），%g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法），%0nd 用于规定输出长度为n的整数，其中开头的数字 0 是必须的。
// %n.mg 用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 或者 f
package c4

import (
	"fmt"
	"time"
)

func init()  {
	fmt.Print(time.Now().Date())
	switch time.Now().Weekday() {
		case time.Sunday, time.Saturday:
			fmt.Print("\n周末放假啦！")
		default:
			fmt.Print("\n工作日。。。")
	}

	t := time.Now()
	switch {
		case t.Hour() < 12: 
			fmt.Print("\n现在是上午！")
		case t.Hour() >= 12: 
			fmt.Print("\n现在是下午！")
	}


	whatAmI := func (i ...interface{}) {
		for index, value := range i {
			
			switch value.(type) {
			case bool:
				fmt.Println("布尔值", index)
			case int:
				fmt.Println("整型",index)
			case string:
				fmt.Println("字符串",index)
			case rune:
				fmt.Println("rune",index)


			}
			
		}
	}
	whatAmI(true, 1, 'a', "abc")

}
