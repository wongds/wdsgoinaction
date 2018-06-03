package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"wdsgoinaction/functional/fib"
)

func tryDefer() {
	//defer里面相当于有一个栈，先进后出
	//只要加了defer就不怕中间return之类的突然出错
	defer fmt.Println(1)
	defer fmt.Println(2)
	panic("error occurr")
	fmt.Println(3)
}

func writeFile(filename string){
	//1 按照文件名创建文件,0666是权限可读可写
	//os.O_EXCL表示创建文件，os.O_EXCL表示如果创建文件时文件存在返回错误信息
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//自己建立错误
	err = errors.New("this is a custom error")
	if err != nil {
		//panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			//不知道错误是什么类型，直接panic
			panic(err)
		} else {
			//知道错误是什么类型直接打印
			fmt.Printf("%s, %s, %s\n", pathError.Op,
					pathError.Path,
					pathError.Err)
		}
		//fmt.Println("file already exists")
		//fmt.Println("Error:", err.Error())

		return
	}
	//2 直接想着关闭文件
	defer file.Close()
	//3 写文件直接用file会比较慢，因此用bufio比较快，包装一下
	//先写到内存里面，大了之后一起倒入过去
	writer := bufio.NewWriter(file)
	//4 写入bufio之后要导入到文件中
	//注意写defer表示最后统一写
	defer writer.Flush()


	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		//fmt.Fprintln(file, f())
		//将后边的print到文件中
		fmt.Fprintln(writer, f())
	}

}

func main() {
	writeFile("fib.txt")
}