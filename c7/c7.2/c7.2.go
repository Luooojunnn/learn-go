// 切片
// 切片是对数组的一个连续片段的引用，也就是引用类型(可以理解为 寄生到数组上)
// 所以在函数之类的想要传递数组的地方，更常用的方式是使用切片（slice），节省空间
// 简言之：切片是一个 长度可变的数组。

package c7

import (
	"fmt"
)

func init() {
	// 0 <= len(s) <= cap(s)

	// 声明方式1：该方式声明出了一个 切片 ，并且隐形声明出了一个长度为3的数组（理解寄生）
	s := make([]string, 3)
	fmt.Println(s)
	
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("s:", s)
	fmt.Println("s[1]:", s[1])
	fmt.Println("len(s):", len(s))

	

	// 声明方式2: var slice1 []type = arr1[start:end]
	var arr = [3]int{1,2,3}
	s2 := arr[:]
	s3 := &arr
	fmt.Println(s2)
	fmt.Println(s3)

	// new() 和 make() 的区别
	// new 函数分配内存，make 函数初始化



	// bytes 包
	// bytes 包也是切片的一种


	// append 切片的追加
	s = append(s, "d")
	s = append(s, "e", "f")
	s = append(s, []string{"z","z","z"}...)
	fmt.Println("追加后的s为：",s)


	// copy 切片复制
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy完后的c：",c)

	l := s[:5]
	fmt.Println("l是s[:5]",l)

	l = s[2:]
	fmt.Println("l是s[2:]",l)

 twoD := make([][]int, 3)
 for i := 0; i < 3; i++ {
	innerLen := i + 1
	twoD[i] = make([]int, innerLen)
	for j := 0; j < innerLen; j++ {
		twoD[i][j] = i + j
	}
 }
 fmt.Println("2d: ", twoD)


}