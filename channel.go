package main

import (
	"fmt"
	"sync"
)

func main() {
	//go Go()
	//time.Sleep(2*time.Second)
	//c := make(chan bool)
	//c := make(chan bool,1)
	//go func() {
	//	fmt.Println("GO!")
	//	c <- true
	//	//<-c
	//	//close(c)
	//}()
	//<-c
	//c <- true
	//for v := range c {
	//	fmt.Println(v)
	//}

	//runtime.GOMAXPROCS(runtime.NumCPU())
	//for i := 0; i < 10; i++ {
		//go channlcache(c, i)
		//fmt.Println(i)
	//}
	//<-c
	//for i := 0; i < 10; i++ {
	//	//c <- true
	//	<-c
	//}

	//设置的CAHNNAL为10
	var c10 chan bool = make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go ca10(c10, i)
	}
	for i := 0; i < 10; i++ {
		<-c10
	}

	//设置WaitGroup,次数为10
	wg:=sync.WaitGroup{}
	wg.Add(10)
	for i:=0;i<10;i++{
		go Wgca(&wg,i)
	}
	wg.Wait()
}

/*
func Go()  {
	fmt.Println("GO!!")
}
*/

func channlcache(c chan bool, index int) {
	a := 1
	for i := 0; i < 100; i++ {
		//fmt.Println(a)
		a += i
	}
	fmt.Println(index, a)
	//c <- true
	//如果INDEX为9  那就channl设置值
	if index == 9 {
		c <- true
		//<-c
	}
}

/**
设置缓存为10的CAHNNAL，执行取值10次
**/
func ca10(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000; i++ {
		//fmt.Println(a)
		a += i
	}
	fmt.Println(index, a)
	c <- true

}

/*WaitGroup: 通过同步包 使用阻塞等待任务完成*/
func Wgca(wg *sync.WaitGroup,index int)  {
	a:=1
	for i:=0;i<10000;i++{
		a+=i
	}
	fmt.Println(index,a)
	wg.Done()
}
