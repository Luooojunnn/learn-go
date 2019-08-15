// 11.1 接口是什么
// 接口是方法签名的命名集合
package c11

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func init()  {
	r := rect{10, 10}
	c := circle{10}

	// circle和rect结构类型都实现了几何接口，因此我们可以使用这些结构的实例作为参数进行测量。

	measure(r)
	measure(c)
}

// 要在go中实现接口，只需要实现接口中的所有方法即可
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 在圆中去实现接口
func (c circle) area() float64 {
	return math.Pi * c.radius *c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量是接口类型，那么它就可以使用接口里的方法
// 如下是一个通用工具函数
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}