package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done" //5
}

func otherTask() {
	fmt.Println("working on something else") //1
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.") //4
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	//retCh := make(chan string)
	//working on something else
	//returned result.
	//Task is done.
	//Done
	//service exited.
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.") //2
		retCh <- ret
		fmt.Println("service exited.") //3
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}


