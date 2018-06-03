package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	aa = 3//这个是包内变量，不是全局变量
	ss = "包内全局变量"
)

func variableZeroValue() {
	var a int
	var s string
	//下面使用printf出带格式的，%q是带引号的，%s是不带引号的
	fmt.Printf("%d %q\n", a, s)//go定义完就有初始值，c定义完int不确定，javanull
}
func variableInitValue() {
	var a, b int = 3, 4
	var s string = "赋初值"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "同行多类型变量声明"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "使用短变量声明,编译器自动决定类型"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler(){
	//float类型小数点后三位，因此.3就能取到0
	fmt.Printf("%.3f\n",cmplx.Exp(1i * math.Pi) + 1)
	//下面实际上结果应该是0
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)//i会以为是变量，1i就知道是叙述
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

func triangle() {
	var a, b int = 3, 4
	//var c int
	//这里类型转换是强制的，go语言不会隐式转换，因此需要两次强制转换，但是注意可能出现4.9这种结果，因为浮点数精度只到小数点后3位
	//c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(calcTriangle(a, b))
}

//常量
func consts() {
	const (//表示一组常量，节省代码量，尽量不要大写，大写代表public
		filename = "abc.txt"
		a, b = 3, 4//常量定义不需要声明，直接=值
	)
	var c int
	//使用常量时，只要实际类型是对的就行，不需要强制转换
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
}

//枚举类型常量
func enums() {
	const(//常量必须定义值
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
	const(
		//使用iota自增值，从0开始
		cpp1 = iota
		_//跳过此值
		python1
		golang1
		javascript
	)
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(cpp1,  python1, golang1, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}


func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss)
	euler()
	triangle()
	consts()
	enums()
}
