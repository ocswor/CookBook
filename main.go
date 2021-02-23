package main

import (
	one "Cookbook/one"
	"Cookbook/persion"
	"fmt"
	"github.com/ocswor/go_utils"
	"os/exec"
)

func demo2(){
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)
	cmd = exec.Command("/bin/bash","-c","ls -l")

	if output,err = cmd.CombinedOutput(); err!=nil{
		fmt.Println(err)
	}

	fmt.Println(string(output))

}
func main() {
	fmt.Println("test")
	//http.ListenAndServe("127.0.0.1:80",handler);
	//b := go_utils.add(1)
	a:=one.Add(2)

	b:=one.TestA(2)
	fmt.Print(a)
	fmt.Print(b)

	x := go_utils.Add(4)
	fmt.Print(x)
	demo2()

	fmt.Println(persion.GregorianToPersian(2020,1,13))
	fmt.Println(persion.PersianToGregorian(1398,10,1))
	fmt.Println(persion.PersianToGregorian(1398,1,1))
	fmt.Println(persion.PersianToGregorian(1403,2,17))
}