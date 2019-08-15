// 错误
// Go 语言的函数经常使用两个返回值来表示执行是否成功：返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败（第 4.4 节）
// 通常（错误信息）都会有像 “Error:” 这样的前缀，所以你的错误信息不要以大写字母开头。
package c5

import (
	"fmt"
	"errors"
)

func init()  {
	for _, value := range []int{7, 42} {
		if r, e := f1(value); e != nil {
			fmt.Println("f1 错误：", e)
		} else {
			fmt.Println("f1 正确：", r)
		}
	}

	for _, value := range []int{7,42} {
		if r, e := f2(value); e != nil {
			fmt.Println("f2 错误：", e)
		} else {
			fmt.Println("f2 正确：", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
    fmt.Println(ae.prob)
	}
}


// 按照惯例，错误是最后一个返回值，并且具有类型错误，即内置接口。
// 一般无错 error 为 nil， 否则把错误抛出来

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("不能传入42")
	}
	return arg+3, nil
}

// 自定义结构类型，显式表示错误
type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	fmt.Println("我被调用了")
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, prob: "不能传入它"}
	}
	return arg+3, nil
}