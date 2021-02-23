package main

import "fmt"

/*
知识点

1.数组类型的值（以下简称数组）的长度是固定的，而切片类型的值（以下简称切片）是可变长的。

2.我们其实可以把切片看做是对数组的一层简单的封装，因为在每个切片的底层数据结构中，一定会包含一个数组。数组可以被叫做切片的底层数组，而切片也可以被看作是对数组的某个连续片段的引用。
也正因为如此，Go 语言的切片类型属于引用类型，同属引用类型的还有字典类型、通道类型、函数类型等；而 Go 语言的数组类型则属于值类型，同属值类型的有基础数据类型以及结构体类型。

3.数组长度 和 容量的区别

4.在底层数组不变的情况下，切片代表的窗口可以向右扩展，直至其底层数组的末尾。切片容量会有一个变化.  base_two 中展示

 */

/*
数据定义

 */



func printArray(arr[5]int)  {
	/*
	参数支持5个元素的素组，就不能传递3个元素的数组。
	在go语言中，这是两种类型。
	 */
	for _,v := range (arr){
		fmt.Println(v)
	}
}

func base_one()  {
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
}

func base_two()  {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
}

func base_three()  {
	s8 := make([]int, 10)
	fmt.Printf("The capacity of s8: %d\n", cap(s8))
	s8a := append(s8, make([]int, 11)...)
	fmt.Printf("s8a: len: %d, cap: %d\n", len(s8a), cap(s8a))
	s8b := append(s8a, make([]int, 23)...)
	fmt.Printf("s8b: len: %d, cap: %d\n", len(s8b), cap(s8b))
	s8c := append(s8b, make([]int, 45)...)
	fmt.Printf("s8c: len: %d, cap: %d\n", len(s8c), cap(s8c))
}
func main() {
	var arr1 [5]int
	arr2:= [3]int{1,2,3}
	var arr3 = [...]int{1,2,3,4,5}

	var arr4 [4][5]int // 4行5列的二维数组
	fmt.Println(arr1,arr2,arr3)
	fmt.Println(arr4)

	//遍历数组
	for i:=0;i< len(arr3);i++{
		fmt.Println(arr3[i])
	}
	fmt.Println("**10")
	for i := range (arr3){
		fmt.Println(arr3[i])
	}

	for i,v := range (arr3){
		fmt.Println(i,v)
	}
	for _,v := range (arr3){
		fmt.Println(v)
	}

	printArray(arr3) //注意 这里数组作为参数传递，是值类型，函数内部并不能改变原始数组的内容。
	// 值传递 是会拷贝数组，考虑优化内存要注意。大的数组 不建议用值传递。



	//base_two()
	base_three()
}
