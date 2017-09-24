package main

import "fmt"

func main()  {

	//数组定义方式
	var array1 [4]string =[4]string{"1","2"}  //前后的4必须加上，不然会报错

	fmt.Println(array1)

	//====================================

	//1.切片的两种定义方式
   var ip = []string{"a","b"," c"}

   var ip1 []string =[]string{"1","2"}
   ip1 = []string{"3","4"}

   fmt.Println(ip)
   fmt.Println(ip1)

   //======make函数定义切片==========
   var ips = make([]string,100)  //100是默认长度，如果用普通的方式定义的话要写很长的代码
   //old方式定义
   //var ips1 = []string{.........},{}里面要写100个
}
