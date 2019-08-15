// 4.2 Go 程序的基本结构和要素
package c4

import (
	"fmt"
)

// 除了符号 _，包中所有代码对象的标识符必须是唯一的
// 大写字母开头，这被称为导出（外部包可用）
// 小写字母开头，整个包的内部是可见并且可用的
// 你可以在使用 import 导入包之后定义或声明 0 个或多个常量（const）、变量（var）和类型（type），这些对象的作用域都是全局的（在本包范围内），所以可以被本包中所有的函数调用

func HelloWorld() {
	fmt.Println("hello world")
}

func init() {
	fmt.Println("我是 c4.2 的 init 函数")
}


// 使用 type 关键字可以定义你自己的类型，你可能想要定义一个结构体(第 10 章)，但是也可以定义一个已经存在的类型的别名，如：
type IZ int
type (
	IZZ int
	FZ float64
	STR string
)
// 这里并不是真正意义上的别名，因为使用这种方法定义之后的类型可以拥有更多的特性，且在类型转换时必须显式转换。
var a IZ = 5


// 类型转换
// 类型 B 的值 = 类型 B(类型 A 的值)
