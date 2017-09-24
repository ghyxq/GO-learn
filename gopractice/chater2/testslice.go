package main

import "fmt"

func main()  {

	//数组定义方式
	var array1 [4]string =[4]string{"1","2"}  //前后的4必须加上，不然会报错

	fmt.Println(array1)

	//====================================

	//切片的两种定义方式
   var ip = []string{"a","b","c"}

   var ip1 []string =[]string{"1","2"}
   ip1 = []string{"3","4"}

   fmt.Println(ip)
   fmt.Println(ip1)
}
