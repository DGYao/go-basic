package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 20 {
			break
		}
		sum += i
	}
	fmt.Println(sum)

	//range方式 相当于java foreach,不过区别在于go能拿到index
	arr := []int{1, 2}
	for v, index := range arr {
		fmt.Println(v, index)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	mArr := make([]int, 0, 5)
	mArr = append(mArr, 1, 2, 3, 4, 5, 6, 7, 8)
	for v, index := range mArr {
		fmt.Print(v, index)
		fmt.Print(",")
	}
	fmt.Println()
	mMap := make(map[string]int, 1)
	mMap["tom"] = 1
	fmt.Println(len(mMap))
	for k, v := range mMap {
		fmt.Print(k, v)
		fmt.Print(",")
	}
}
