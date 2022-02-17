package main

import "fmt"


// go 语言方法和接口的使用
// go 语言种Interface 是可能为nil 空指针，所以针对interface 一定要预先判空，否则会引起程序crash（nil panic）
// Struct 初始化意味着空间分配，对struct 引用不会出现空指针

type Interface interface{
	getName() string
}

type Human struct{
	firstName, lastName string
}

// 接口实现不需声明，对Human 对象添加方法
// 指针类型实现，接口找不到上实现类
// func (h *Human) getName() string{
// 	return h.firstName + "," + h.lastName
// }


func (h Human) getName() string{
	return h.firstName + "," + h.lastName
}

type Car struct{
	factory, model string
}

func (c Car) getName() string{
	return "品牌" + c.factory + "车型" + c.model
}

func main(){
	interfaces := []Interface{}
	human := new(Human)
	interfaces = append(interfaces, human)
	human.firstName = "Marvin"
	human.lastName = "Yang"
	
	car := new(Car)
	car.factory = "JiangSu"
	car.model = "Dongfeng"

	interfaces = append(interfaces, car)
	for _, f := range interfaces {
		fmt.Println(f.getName())
		// f.(Car) 接口独有，断言
		// fmt.Println(f.(Car))

		switch value := f.(type) {
			case *Car:
				fmt.Println("this is a Car Object")			
			case *Human:
				fmt.Println("this is a Human Object")
			default:
				fmt.Printf("default %T" , value)
		}
	}

	// 上下比较 interface 定义和非定义判断方法也有不同
	mapInterfaces := make(map[string]interface{}, 2)
	// 这里一定要判断 定义的变量是什么类型，类型错误，处理不正确
	// Human{} type 为Human | new(Human) 类型为指针 *Human
	mapInterfaces["human"] = Human{}
	mapInterfaces["car"] = car
	mapInterfaces["string"] = "string"

	fmt.Println(mapInterfaces)

	for key, f := range mapInterfaces {
		fmt.Println(key)
		switch value := f.(type) {
			case string:
				fmt.Println("this is a string")
			case Human:
				fmt.Println("this is a Human Object")
			case interface{}:
				fmt.Println("this is a Car Object")			
			default:
				fmt.Printf("default %T" , value)
		}
	}

}


