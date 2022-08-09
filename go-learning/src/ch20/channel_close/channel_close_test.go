package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		//panic: send on closed channel
		//ch <- 11
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		//程序还是正常退出，并没有被阻塞在这个地方，因为我们关闭了通道。
		//第11次我们想要从通道里面接收数的时候，因为通道已经关闭，
		//虽然我们没有判断它是不是被关闭，所以说他会迅速返回通道所定义类型的零值
		//for i := 0; i < 11; i++ {
		//	data := <-ch
		//	fmt.Println(data)
		//}
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	wg.Wait()
}
