package main

import (
	"fmt"
)


//接口demo
type Phone interface {
	Call()
	Receive() string
	Getsumsung() Sumsung
}

type Iphone struct {
	Message string
}

type Sumsung string

func (iphone Iphone) Call() {

	fmt.Println("这是iphone！")
}

func (iphone Iphone) Receive() string {

	return iphone.Message
}

func (iphone Iphone) Getsumsung() Sumsung {

	return Sumsung("iphone实现getsumsung()")
}

func (s Sumsung) Call() {

	fmt.Println("这是Sumsung！")
}

func (s Sumsung) Receive() string {

	return "Sumsung实现receive()"
}

func (s Sumsung) Getsumsung() Sumsung {
	return s
}

func main() {
	fmt.Println("=================")

	//接口demo
	var Iphone1 Iphone
	Iphone1.Message = "你好，我是iphone"
	name := Sumsung("你好，我是sumsung")
	var phone Phone
	var phone01 Phone
	phone = Iphone1
	phone01 = name
	fmt.Println(phone.Receive())
	fmt.Println(phone.Getsumsung())
	phone.Call()
	fmt.Println(phone01.Receive())
	fmt.Println(phone01.Getsumsung())
	phone01.Call()


}