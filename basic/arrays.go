package main


import "fmt"

func printArray(arr *[5]int){
	//这里调用的方式非常方便，即使是指针类型也能像原来的值类型一样直接调用
	arr[0] = 100
	for i := 0; i < len(arr); i++{
		fmt.Println(arr[i])
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(grid)
	fmt.Println(arr1, arr2, arr3)

	//遍历
	for i := 0; i < len(arr3); i++{
		fmt.Println(arr3[i])
	}
	//实际使用的方法,同时获得index和元素的值
	for i, v := range arr3{
		fmt.Println(i, v)
	}
	//只想要值
	for _, v := range arr3{
		fmt.Println(v)
	}
	//数组是值类型,使用指针才能改变原来的值
	fmt.Println("开始打印arr1")
	printArray(&arr1)
	fmt.Println("调用函数打印arr3")
	printArray(&arr3)
	fmt.Println("开始打印arr3")
	fmt.Println(arr3)

	//要想改变原来的值，还有一种方法，就是将传入的参数改成slice
	//然后传入的参数arr数组要想改成slice，只需要传递的时候写成arr[:]就行了
}