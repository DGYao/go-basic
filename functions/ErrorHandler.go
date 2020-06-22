package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, e := os.Open("f.txt")
	fmt.Println(file, e)
	e2 := errorHandler("aaa")
	if e2 != nil {
		fmt.Println(e2)
	}
	errorCatch(9)

}

func errorHandler(name string) error {
	file, e := os.Open(name)
	if e != nil {
		return errors.New("找不到文件")
	}
	fmt.Println(file)
	return nil
}

func errorCatch(index int) (a int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error", r)
			a = -1
		}
	}()
	arr := []int{1, 2, 3}
	return arr[index]
}
