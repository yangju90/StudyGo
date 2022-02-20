package main

import(
	"fmt"
	"errors"
)

type ErrorStatus struct{
	msg string
}

func (e ErrorStatus) Error() string{
	return e.msg
}

type Test interface{
	test() string
}

// 闭包
func Inc() (v int) { 
   
    defer func(){ 
        v++ 
    } () 
 
    return 42 
}

// try catch 语句
func panicRecover() {
	defer func() {
		fmt.Println("defer func is called!")
		if err := recover(); err != nil{
			fmt.Println("Recovering ......")
			fmt.Println(err)
		}
	}()

	panic("a panic is triggered !")
}

func main(){

	panicRecover()

	fmtError := fmt.Errorf("fmt package error Testing!")
	fmt.Println(fmtError)

    structError := ErrorStatus{"struct implements interface error testing!"}
	fmt.Println(structError.Error())

	interfaces := []error{}

	interfaces = append(interfaces, structError)

	fmt.Println(interfaces[0].Error())

	value,ok := interfaces[0].(Test)
	fmt.Println(value, ok)

	number := 10
	newNumber := Inc()

	fmt.Println(number)
	fmt.Println(newNumber)

	var errNotFound error = errors.New("error Not Found!")

	fmt.Println(errNotFound)
	
}