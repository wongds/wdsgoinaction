package main

import (
	"fmt"
)

func updateSlice (s []int){
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := arr[:]
	s3 := arr[2:]
	//注意任何时候都是左开右闭
	fmt.Println("arr[2:6] = ", s1)
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("s3 = ", s3)
	fmt.Println("s2 = ", s2)
	//lice是对原来数组的一个view视图
	fmt.Println("After updateSlice(s3)")
	updateSlice(s3)
	fmt.Println("s3 = ", s3)
	fmt.Println("arr = ",arr)


	fmt.Println("Reslice通过slice构造slice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr=", arr)
	s1 = arr[2:6]
	s2 = s1[3:5]
	//下面说明slice可以向后扩展，cap是底层数组的容量
	fmt.Printf("s1 = %v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	ss3 := append(s2, 10)
	//s4和s5view不再是arr了，而是一个新的数组。
	//添加元素是如果超越cap，系统会重新分配更大的底层数组
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5=", ss3, s4, s5)
	fmt.Println("arr = ", arr)
}
