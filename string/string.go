package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "GoLand"
	str2 := "Go语言"
	fmt.Println(str1[2])
	fmt.Println(str2[2])
	fmt.Println(string(str2[2]))
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	fmt.Println(len(str2))

	runeArr := []rune(str2)
	fmt.Println(string(runeArr[2]))
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())
	fmt.Println(len(runeArr))
}
