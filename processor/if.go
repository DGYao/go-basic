package main

import (
	"fmt"
)

func main() {
	if age := 10; age < 18 {
		fmt.Println("未成年")
	} else {
		fmt.Println("成年")
	}
}
