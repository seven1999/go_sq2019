package cal

import (
"fmt"
)

// MyIterface

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
