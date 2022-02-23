package main

import (
	"fmt"
	"reflect"
)

func main() {
	//某结构的基本操作
	u := User{1, "ppp", 12}
	//指针类型，必须使用指针的反射方法
	//Info(&u)
	Info(u)

	//取匿名函数
	m := Manager{User{2, "ooo", 13}, "nbl"}
	t := reflect.TypeOf(m)
	//取出第一个反射的字段,Anonymous:true
	fmt.Printf("%#v\n", t.Field(0))
	//匿名字段的值 FieldByIndex([]int{0,1})="Name" | FieldByIndex([]int{0,2})="Age"
	fmt.Printf("%#v\n",t.FieldByIndex([]int{0,0}))

	//反射修改类型
	x:=123
	//必须使用指针引用对象，才能修改
	v:=reflect.ValueOf(&x)
	v.Elem().SetInt(99)
	fmt.Println(x)

	//反射修改结构 set()
	Set(&u)
	fmt.Println(u)

	//反射动态调用方法,传入一个参数,u.Hello()
	u3:=User{3,"yy",15}
	u3.Hello("ffff")
	v3:=reflect.ValueOf(u3)
	//select func: Hi()
	mv:=v3.MethodByName("Hello")
	//反射调用方法传参执行,args is slice.
	args:= []reflect.Value{reflect.ValueOf("zy")}
	mv.Call(args)
}

type User struct {
	Id   int
	Name string
	Age  int
}

//匿名字段
type Manager struct {
	User
	title string
}

func (u User) Hi() {
	fmt.Println("hello,", u.Name,"my years:",u.Age)
}

func (u User) Hello(name string) {
	fmt.Println("hello,", name,"my name is ",u.Name)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t)

	v := reflect.ValueOf(o)
	fmt.Println("Fieds:", v)

	//反射出结构字段信息
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v=%v\n", f.Name, f.Type, val)
	}
	//反射出结构方法信息
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

//反射修改结构, 必须是可以修改的, 且指针调用
func Set(o interface{}){
	//取出当前实例的值
	v:=reflect.ValueOf(o)
	//判断类型为指针 且 可修改, else给出值
	if v.Kind()==reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("xxxx")
	}else {
		v=v.Elem()
	}
	//判断 是否能取得当前字段,没有找到 为空vlaue
	//f:=v.FieldByName("Name1")
	f:=v.FieldByName("Name1")
	if !f.IsValid(){
		fmt.Println("nil value!!")
		return
	}
	// 使用字段名字取值
	if f:=v.FieldByName("Name");f.Kind()==reflect.String{
		f.SetString("changed")
	}

}
