package main


import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		q, _ := div(a, b)
		return q, nil
	default:
		//panic会中断执行
		//panic("unsupported operation: " + op)
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div (a, b int) (q, r int){
	return a / b, a % b
}

func apply (op func(int, int) int, a, b int) int {
	//通过反射获得函数指针
	p := reflect.ValueOf(op).Pointer()
	//获得函数名
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" +
		"(%d, %d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

//值传递不能交换值，但是指针传递能
func swap(a, b *int){
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int){
	return b, a
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(fmt.Println(result))
	}
	fmt.Println(eval(3, 4, "*"))
	q, r := div(13, 3)
	fmt.Println(q, r)
	//传入函数及其参数，返回函数结果，并通过反射获得函数名
	fmt.Println(apply(pow, 3, 4))
	//直接定义匿名函数传入参数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
	fmt.Println(sum(1, 2, 3))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println(swap1(a, b))
}
