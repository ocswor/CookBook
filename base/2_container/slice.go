package main

import "fmt"

func updateSlice(s[]int)  {
	s[0] = 100
}
func main2() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6] = ",arr[2:6])
	fmt.Println("arr[:6] = ",arr[:6])
	fmt.Println("arr[2:] = ",arr[2:])
	s1 := arr[:]  // s1 就是slice 内部不存储数据，是arr的view
	fmt.Println("arr[:]=",s1)
	updateSlice(s1)

	fmt.Println("arr[:]=",s1)
	fmt.Println("arr=",arr)

	s2 := arr[2:]
	updateSlice(s2)
	fmt.Println("arr[:]=",s2)
	fmt.Println("arr=",arr)
}