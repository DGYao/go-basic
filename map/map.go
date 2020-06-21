package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m2 := map[string]int{
		"tom": 1,
		"yao": 2,
	}
	m1["tom"] = 2
	fmt.Println(m1, m2)

}
