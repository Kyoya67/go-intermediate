package main

import "fmt"

type Dog struct{}

func (d Dog) Speak() string {
	return "wan"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "myao"
}

type Speaker interface {
	Speak() string
}

func main() {
	var hello interface{} = "hello"
	var dog interface{ Speak() string } = Dog{}
	var cat interface{} = Cat{}

	s, ok := hello.(string)
	fmt.Println(s, ok)

	d1, ok := dog.(Dog)
	fmt.Println(d1.Speak(), ok)

	d2, ok := dog.(Cat)
	fmt.Println(d2.Speak(), ok)

	c1, ok := cat.(Speaker)
	fmt.Println(c1.Speak(), ok)

	c2, ok := cat.(Cat)
	fmt.Println(c2.Speak(), ok)

	//catの実態はCat{}だけど、型アサーションしないとSpeakメソッドは使えないためコンパイルエラーになる
	//fmt.Println(cat.Speak())
}
