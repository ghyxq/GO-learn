package main

import "fmt"

//_,ok := talk.(&Mytalk) 这句话还有问题


//接口定义
type talk interface {
	Hello()
	World()
}


//实现上面的接口，下面33行用了类型断言
type Mytalk string  //mytalk实现了talk类型

func (m *Mytalk) Hello()  {

}

func (m *Mytalk) World()  {

}

func main()  {

	//var talktest talk= Mytalk("ok")  如果Mytalk实现接口则OK，如果是*Mytalk，则有问题

	s := Mytalk("ok")
	var talktest talk= &s    //这里不能直接写var talktest talk = &Mytalk("ok") 会报找不到 cannot take the address of Mytalk("hello")

	//类型断言，判断talktest是否实现了talk接口
	ok := talktest.(talk)
	fmt.Println("hello world",ok)


}