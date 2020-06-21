package main

import "fmt"

func main() {
	type Sex int8
	const (
		MALE   Sex = 1
		FEMALE Sex = 2
	)
	sex := MALE
	//匹配到默认不会往下执行
	//想继续往下执行 加fallthrough
	switch sex {
	case MALE:
		fmt.Println("male")
		fallthrough
	case FEMALE:
		fmt.Println("female")
	default:
		fmt.Println("unKnow sex")
	}
}
