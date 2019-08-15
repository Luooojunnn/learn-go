// 4.4 变量
package c4

import (
	"fmt"
	"os"
)

// 它在声明变量时将变量的类型放在变量的名称之后
var c = 1
var a, b *int = &c, &c

// 这种因式分解关键字的写法一般用于声明全局变量。
var (
	str1 string
	num1 int
	bol1 bool
)

// 变量的类型也可以在运行时实现自动推断
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("PATH")
)

// 4.4.2 值类型和引用类型

// 所有像 int、float、bool 和 string 这些基本类型都属于值类型
// 另外，像数组（第 7 章）和结构（第 10 章）这些复合类型也是值类型
// 值类型存储在栈中
// 指针（第 4.9 节）属于引用类型，其它的引用类型还包括 slices（第 7 章），maps（第 8 章）和 channel（第 13 章）
// 引用类型存储在堆中，比栈拥有更大的内存空间。

// 4.4.3 打印
// Printf 可以在 fmt 包外部使用
// %s 代表字符串标识符
// %v 代表使用类型的默认输出格式的标识符
// 函数 fmt.Print 和 fmt.Println 会自动使用格式化标识符 %v 对字符串进行格式化

// 4.4.5 init 函数
// 它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高
// 一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性
// init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine

// 练习
const a1 = 'a'
func show(){
	n()
	m()
	n()
}
func n(){
	fmt.Print(a1)
}
func m(){
	a1 := 0
	fmt.Print(a1)
}

func init(){
	show()
}

