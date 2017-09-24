package main

import "fmt"


type talk interface {
	Hello()
	World()
}



type Mytalk string  //mytalk实现了talk类型


func (m *Mytalk) Hello()  {

}

func (m *Mytalk) World()  {

}

func main()  {

	//var talktest talk= Mytalk("ok")  如果Mytalk实现接口则OK，如果是*Mytalk，则有问题

	s := Mytalk("ok")
	var talktest talk= &s    //这里不能直接写var talktest talk = &Mytalk("ok") 会报找不到 cannot take the address of Mytalk("hello")

	ok := talktest.(talk)
	fmt.Println("hello world",ok)


}