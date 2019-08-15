//  结构（struct）
// 可以简单理解为：结构就是类式语言里边的实体对象
// 因此可以通过 new 函数来创建。

package c10

import (
	"fmt"
)

type persopn struct {
	name string
	age int
}

func init()  {
	fmt.Println(persopn{"luojun",20})

	fmt.Println(persopn{name: "zhangsan", age: 100})

	fmt.Println(persopn{name: "lisi"})

	// 获得实例的指针
	fmt.Println(&persopn{name: "wanger", age: 10})

	// . 获取
	p := persopn{name: "李牛", age: 23}
	fmt.Println(p.name)

	// 也可以使用指针
	s := &p
	fmt.Println(s.name)

	// 设置
	p.name = "李牛1"
	fmt.Println(p.name)
	s.name = "李牛2"
	fmt.Println(s.name)
}