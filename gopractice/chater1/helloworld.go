package main

import (
	"bufio"
	"os"
	"fmt"
)

func main()  {

	//获取标准输入
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("please input your name:")
	input,err := inputReader.ReadString('\n')

	if err != nil{
		fmt.Println("there is a error",err)
	}else{

		input = input[:len(input)-1]

		fmt.Println("hello ",input)
	}


}
