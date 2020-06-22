package main

import "fmt"

func main() {
	i := add(1, 2)
	fmt.Println(i)
	i2, i3 := div(4, 2)
	fmt.Println(i2, i3)
	a, b := name()
	fmt.Println(a, b)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num2 % num1
}

func name() (a int, b int) {
	a = 1
	b = 2
	return
}
