// range 
// 如何迭代查询一些结构的元素呢
package c7

import (
	"fmt"
)

func init() {
	// 遍历切片(数组是一样的)
	sli1 := []int{1,2,3,4,5}
	sum := 0
	for _, v := range sli1 {
		sum += v
	}
	fmt.Println("sum",sum)
	for i, v := range sli1 {
		if i%2 == 0 {
			fmt.Printf("偶数：%d -- %d\n",i,v)
		}
	}

	// 遍历map结构
	map1 := map[string]string{"name": "luojun", "age": "100"}
	for k, v := range map1 {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// map的话，也可以只要key
	for k := range map1 {
		fmt.Printf("key: %s\n", k)
	}

	// 遍历字符串也是可以的
	for i, v := range "abdefg" {
		fmt.Println(i, v)
	}
}