
/**
包中的init函数初始化，在包全局变量初始化之后，在main函数执行之前

*/

package main

import "fmt"


//该函数会在该包全局变量初始化后调用
func init()  {

	fmt.Println(hello)
	fmt.Printf("MAP: %v",m)

}


var hello = "hello world"    //全局变量先初始化，然后在init函数初始化
var m = map[int]string{1:"a",2:"b"}


var info string

func main()  {

	fmt.Println(hello)

}
