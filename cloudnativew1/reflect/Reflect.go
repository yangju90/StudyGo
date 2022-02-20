package main

import (
	"reflect"
	"fmt"
)

// go 语言反射调用

type Person struct{
	name string
	sex string
	//Alias是公有字段定义方式（反射可修改它的值），age是私有字段定义方式（反射不可修改它的值）
	Alias string
}

// go语言反射 只能拿到 本身类型的反射方法 拿不到指针类型的反射方法
// 仅能获取public方法
func (p Person) GetName() string {
	return p.name
}

func (p *Person) GetSex() string {
	return p.sex
}


func main(){
	myMap := make(map[string]string,  10)
	myMap["a"] = "b"
	myMap["c"] = "d"

	// 反射获取类型
	mapType := reflect.TypeOf(myMap) 
	fmt.Println("myMap type:" ,  mapType)
	// 反射获取值
	mapValue := reflect.ValueOf(myMap)
	fmt.Println("myMap value:" , mapValue)

	// struct
	myStruct := Person{name:"Marvin", sex:"man", Alias:"bbb"}

	v1 := reflect.ValueOf(myStruct)

	for i := 0; i < v1.NumField(); i++ {
		fmt.Println("Field " , v1.Field(i))
	}

	fmt.Println(v1.CanSet())
	v2 := reflect.ValueOf(&myStruct)
	fmt.Println(v2)
	v := v2.Elem();
	fmt.Println(v.CanSet())
	// 私有变量不可修改 false
	fmt.Println(v.Field(0).CanSet())
	// 私有变量不可修改 false
	fmt.Println(v.Field(1).CanSet())
	// public 可修改 true
	fmt.Println(v.Field(2).CanSet())

	v.Field(2).SetString("bbb")

	fmt.Println(v2)
	fmt.Println(myStruct)

	fmt.Println(myStruct)
	
	

	fmt.Println(v1.NumField())
	fmt.Println(v1.NumMethod())
	fmt.Println(myStruct.GetName())
	for i := 0; i < v1.NumMethod(); i++ {
		fmt.Println("Method " ,  v1.Method(i))
	}



}