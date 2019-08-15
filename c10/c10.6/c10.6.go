// 方法！
// 可以简单理解为：结构是类，给类的实例对象添加方法
package c10

import (
	"fmt"
)

type rect struct {
	width, height int
}

func init() {
	r := rect{width: 10, height: 5}

	fmt.Println("r指针函数",r.area())
	fmt.Println("r非指针函数",r.perim())

	rp := &r
	fmt.Println("rp指针函数",rp.area())
	fmt.Println("rp非指针函数",rp.perim())
}

// 为结构定义方法的时候，可以是指针可以不是

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}