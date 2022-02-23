package main

import (
	"fmt"
	"time"
)

//多个CHANNEL操作
func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	//close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}

	c3 := make(chan int)
	go rp(c3)

	//for{ //卡死事件循环
	for i := 0; i < 5; i++ {
		//select 发送消息
		select {
		case c3 <- 0:
		case c3 <- 1:
		}
	}

	//close(c3)
	timeout()

}

//随机输出(0\1) --c3
func rp(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

//timeout:select configure
func timeout() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout.....")
	}
}
