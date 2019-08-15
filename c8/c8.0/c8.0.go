// 8.0 Map
// map 是一种特殊的数据结构：一种元素对（pair）的无序集合，pair 的一个元素是 key，对应的另一个元素是 value，所以这个结构也称为关联数组或字典。这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位
// key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key (译者注：含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的），但是指针和接口类型可以

package c8

import (
	"fmt"
)


// 不要使用 new，永远用 make 来构造 map !!!

func init()  {
	// 基础声明形式
	mapCreated := make(map[string]int)

	mapCreated["name"] = 987
	mapCreated["age"] = 20

	fmt.Println("mapCreated",mapCreated)

	// 取值
	name1 := mapCreated["name"]
	fmt.Println("name1",name1)

	// 获取长度
	fmt.Println("len",len(mapCreated))

	// map 中移除键值对
	delete(mapCreated, "age")

	fmt.Println("mapCreated",mapCreated)

	// 为了区分，是有这个key，但是值就是nil，还说说没有这个key导致的值是nil，那么就根据第二个返回值来判断
	_, age := mapCreated["age"]
	fmt.Println("age",age)


	// 声明和赋值在一起
	n := map[string]string{"name": "luojun", "age": "18"}
	fmt.Println("n", n)

}

