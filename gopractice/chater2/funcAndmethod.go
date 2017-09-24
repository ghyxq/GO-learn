package main

import (
	"fmt"
	"errors"
)

//1.函数定义
//
func add(a,b int) (int,error)  {

	return a+b,errors.New("ok")//这里还有点问题，该怎么返回
}

func add1(a,b int) (sum int,err error)  { //返回值的名称要么都没有，要么都有，有的话在函数调用的时候会隐式声明这两个变量

	sum= a+b  //这里可以直接使用返回值的变量，
	return
}



func main()  {

	fmt.Println("test func and method")

	add(1,2)
	add1(3,4)
}
