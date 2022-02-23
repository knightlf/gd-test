package main

import "fmt"

func main() {
	//use Phone interface to Onpluse
	a := Onepluse{"one", 13811802222}
	a.call()

	//use Phone interface to IPhone
	b := IPhone{Name: "ip", Sri: "hiiiii"}
	msg := b.sendmsg()
	print(msg)

	b.AppSc()
	Disconnect(b)

	//对象赋值给接口，是拷贝形式，内部存储使用的复制品指针，无法修改状态
	b.Name="nip"
	b.AppSc()
	Disconnect(b)

}

/* 定义接口 */
type Phone interface {
	call()
	sendmsg() string
	Sc
}

/* 子接口 */
type Sc interface {
	AppSc()
}

/* 定义结构体 */
type Onepluse struct {
	Name string
	Numb int
}

type IPhone struct {
	Name string
	Sri  string
}

/* 实现接口方法 */
func (iPhone IPhone) call() {
	fmt.Println("im iphone,call u!")
}

func (one Onepluse) call() {
	fmt.Println("im the one.call me", one.Numb)
	fmt.Println("number:", one.Numb)
	fmt.Println("name:", one.Name)
}

func (iPhone IPhone) AppSc() {
	fmt.Println("im iphone app store connectd.",iPhone.Name)
}

/* 实现接口方法 */
func (iPhone IPhone) sendmsg() string {
	return "iphone send msg"
}

func (one Onepluse) sendmsg() string {
	return "one send msg"
}

/* 断言，类型的判断
传入一个接口 进行 检验
*/
func Disconnect(p interface{}) {
	if pc, ok := p.(IPhone); ok {
		fmt.Println("Disconnect from:", pc.Sri)
	}
	if pc, ok := p.(Phone); ok {
		fmt.Println("Disconnect from:", pc.sendmsg())
	}
	fmt.Println("Unknow dev.")
	/*
		type swich:多类型检测, 不指定类型，使用(type)
	*/
	switch v := p.(type) {
	case Onepluse:
		fmt.Println("S Disconnect from:", v.Name)
	case IPhone:
		fmt.Println("S Disconnect from:", v.Sri,"++++",v.Name)
	case Sc:
		fmt.Println("S Disconnect from:", v.AppSc)
	default:
		fmt.Println("S Unknow dev.")
	}

}
