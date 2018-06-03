package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		}else {
			panic(fmt.Sprintf("我不知道要干嘛: %v", r))
		}
	}()//定义匿名函数时函数体后面直接加括号相当于调用
	//panic(errors.New("this is an error"))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	//这样会报两次，因为defer中也有调用这个函数的
	panic(123)
}

func main() {
	tryRecover()
}