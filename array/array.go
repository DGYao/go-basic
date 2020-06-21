package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	fmt.Println(arr1)
	for i := 0; i < len(arr1); i++ {
		arr1[i] += 1
	}
	fmt.Println(arr1)
	//切片
	mk := make([]int, 1)
	fmt.Println(mk)
	fmt.Println(len(mk), cap(mk)) //长度，容量

	mk = append(mk, 3, 4, 5, 6) //只有未定义的变量才需要:=
	mkSub1 := mk[3:5]
	mkSub2 := mk[5:6]
	fmt.Println(mkSub1)
	fmt.Println(mkSub2)
	combine := append(mkSub2, mkSub1...) //...是将mkSub1解开成一个个元素
	fmt.Println(combine)

}
