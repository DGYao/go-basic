package main

import "fmt"

func main() {
	//指针赋值会改变所有指向同一个地址的值
	str := "GoLand"
	var point = &str
	*point = "go"
	fmt.Println(str)
	i := 1
	pointf(&i, i) //函数传递指针会改变值
	fmt.Println(i)
}

func pointf(point *int, value int) {
	*point = 2
	value = 3
}
