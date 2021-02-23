package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

/*

 */

func base1() {
	var (
		cmd *exec.Cmd
		err error
	)
	cmd = exec.Command("/bin/bash", "-c", "echo 1")

	err = cmd.Run()
	fmt.Println(err)
}

func base2() {
	var (
		cmd    *exec.Cmd
		err    error
		output []byte
	)
	cmd = exec.Command("/bin/bash", "-c", "sleep 2;echo 1")

	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

}

func base3() {

	type result struct {
		err    error
		output []byte
	}
	//var (
	//	ctx context.Context
	//	cancel context.CancelFunc
	//)

	//resultChan := make(chan *result, 1000)
	resultChan := make(chan *result)
	ctx, _ := context.WithCancel(context.TODO())

	go func() {
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c","sleep 1;ls")

		output, err := cmd.CombinedOutput()

		resultChan <- &result{
			err:    err,
			output: output,
		}

	}()

	fmt.Println("x")
	time.Sleep(2 * time.Second)

	fmt.Println("2")
	//cancelFunc()

	res := <-resultChan
	//fmt.Println(res.err, string(res.output))
	fmt.Printf("result error:%v  output:%s",res.err,string(res.output))
}
func main() {

	//base2()

	base3()
}
