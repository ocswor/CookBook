//变量定义 学习

//常量定义
//枚举类型


// if xx; dd{
//}

//go  switch 不用break  支持字符串 跳转

//
package main

import (
	"fmt"
	"math"
)

// 作用域  在包内部的变量
var aa = 3
var ss = "kkk"

var (
	rr = 4
	bb = true
)

func variableZeroValue(){
	// 定义变量会有默认值
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s) //空字符串会显示出来。
	fmt.Printf("%d %s\n",a,s)
}

func variableInitialValue()  {
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

func variableTypeDeduction()  {
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func variableShorter()  {
	a,b,c,s := 3,4,true,"def"
	b = 5
	fmt.Println(a,b,c,s)
}

func consts()  {
	/*
	常量定义不用 声名类型
	编译器内部操作 相当于内部替换
	 */
	const filename  = "abc"
	const a,b = 3,4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename,c)
}

func enums()  {
	/*
	枚举实例
	_ 可以跳过
	 */
	const (
		cpp = iota
		java
		python
		_
		golang
	)
	fmt.Println(cpp,java,python,golang)

	const(
		b = 1<< (10 *iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)

}
func main() {
	variableZeroValue()
 	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,ss,rr)
	enums()
}

