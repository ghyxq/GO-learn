package main

import (
	"bufio"
	"os"
	"fmt"
)

func main()  {

	//读取终端输入
	bufferReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name")

	inputString,err:= bufferReader.ReadString('\n')



	if err !=nil{
		fmt.Println("there has an error")
		fmt.Printf("An error occurred:%s\n",err)

	}else{

		inputString = inputString[:len(inputString)-1]
		fmt.Printf("Welcome %s to here",inputString)
	}

	//循环读取
	for{

		inputString,err = bufferReader.ReadString('\n')

		if err!=nil {
			fmt.Println("an error occurred!")
			fmt.Printf("%s there occurred an error ,run hurry up!",inputString)
			continue

		}

		inputString = inputString[:len(inputString)-1] //去掉最后一个换行符

		switch inputString {

		case "hello":
			fmt.Println("hello ha")
		case "":
			continue
		case "bye":
			fmt.Println("bye bye!")
			os.Exit(0)
		default:
			fmt.Println("sorry,i down understand your meaning,please input other words ")


		}



	}

}
