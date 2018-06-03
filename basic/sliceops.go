package main

import (
	"fmt"
)
//运行发现slice的cap是成倍增加的，从0,1开始
func printSlice(s []int){
	fmt.Printf("%v, len=%d, cap=  %d\n", s, len(s), cap(s))
}

func main(){
	var s []int//Zero value for slice is nil
	for i := 0; i < 100; i++{
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	//创建有len的slice
	s2 := make([]int, 16)
	//创建有len和cap的slice
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)
	//复制元素
	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)
	//删除中间元素
	fmt.Println("Deleting elements from slice")
	//通过覆盖的方式，覆盖了原来8位置的数字，相当于删除了，len减小了1位，但是cap没变
	s2 = append(s2[:3], s2[4:]...)//...表示将对应位置后面全覆盖添加第二个slice的东西
	printSlice(s2)
	//删除开头元素
	fmt.Println("Popping from front")
	//取出开头元素
	front := s2[0]
	//删除开头元素
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)
	fmt.Println("Popping from back")
	//取出尾巴元素
	tail := s2[len(s2) - 1]
	//删除尾巴元素
	s2 = s2[:len(s2) - 1]
	fmt.Println(tail)
	printSlice(s2)


}