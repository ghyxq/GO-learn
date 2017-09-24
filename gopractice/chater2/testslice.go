package main

import "fmt"

func main()  {

	//*******************array****************
	//数组定义方式
	var array1 [4]string =[4]string{"1","2"}  //前后的4必须加上，不然会报错

	fmt.Println(array1)

	//****************slice****************

	//1.切片的两种定义方式
   var ip = []string{"a","b"," c"}

   var ip1 []string =[]string{"1","2"}
   ip1 = []string{"3","4"}

   fmt.Println(ip)
   fmt.Println(ip1)

             //--------------make函数定义切片--------------
   			var ips = make([]string,100)  //100是默认长度，如果用普通的方式定义的话要写很长的代码
  			 //old方式定义
			//var ips1 = []string{.........},{}里面要写100个
	       fmt.Println(ips)

	//******************map==是引用类型****************

	var map1 = map[string]bool{} //key类型是string，value类型是bool，大小为空
	map1["hello"]=true  //赋值

	fmt.Println(map1)
    delete(map1,"hello")//删除map1中的元素，不管map中有没有，都会执行

	fmt.Println(map1)


}
