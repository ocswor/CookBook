package main

import (
	"errors"
	"fmt"
	"log"
)

type Printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

type operate func(x, y int) int
type calculateFunc func(x int, y int) (int, error)

//  什么是闭包
// 在一个函数中存在对外来标识符的引用。所谓的外来标识符，既不代表当前函数的任何参数或结果，也不是函数内部声明的，它是直接从外边拿过来的。
func genCalculator(op operate) calculateFunc {
	// op 自由变量究竟代表了什么，这一点并不是在定义这个闭包函数的时候确定的，而是在genCalculator函数被调用的时候确定的。
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	var(
		p Printer
		x int
		y int
		result int
		err error
	)
	p = printToStd
	if _,err :=p("something"); err != nil {
		log.Print("x")
	}

	// 函数 封装成另外一个函数
	op := func(x, y int) int {
		return x + y
	}

	x, y = 56, 78
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}
