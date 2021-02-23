package main

import (
	"fmt"
	"math/rand"
	"time"
)

func base1() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 2
		ch1 <- 1
		ch1 <- 3
	}()

	//ch1 <- 4
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
}

func base2() {
	/*
		channel接收和发送数据都是阻塞的

		报错 fatal error: all goroutines are asleep - deadlock!
	*/
	ch1 := make(chan int)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	//ch1 <- 4
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
}

func base3() {
	// 示例1。
	ch1 := make(chan int, 1)
	ch1 <- 1
	//ch1 <- 2 // 通道已满，出现异常

	// 示例2。
	fmt.Println("ch1")
	ch2 := make(chan int, 1)
	fmt.Println("ch2x")
	//elem, ok := <-ch2 // 通道已空，出现异常

	//fmt.Println("ch2")
	//_, _ = elem, ok
	ch2 <- 1

	// 示例3。
	var ch3 chan int
	ch3 <- 1 // 通道的值为nil，出现异常
	//<-ch3 // 通道的值为nil，
	_ = ch3

}

func base4() {

	// 准备好几个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}

func base5() {
	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		//close(intChan)
		intChan <- 1
	})
	for{

		select {
		case v, ok := <-intChan:
			fmt.Printf("%v %v\n",ok,v)
			if !ok {
				fmt.Println("The candidate case is closed.")
				break
			}
			fmt.Println("The candidate case is selected.")
		}
	}

}
func main() {
	//base1()
	//base2()
	//base3()

	//base4()

	base5()

}
