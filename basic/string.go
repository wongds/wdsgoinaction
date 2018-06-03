package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"//UTF-8可变长，英文就是一字节，中文是三字节
	fmt.Println(s)
	//[]byte获取字节
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	//这里面i是指字节开始的位置，并不是连续的，因为中文是三个字节的。这里会将rune转换成unicode
	//并不是底层数组存储直接映射，而是解码之后重新存储
	for i, ch := range s {//ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	//len知识获得字节长度
	//直接数字符数
	fmt.Println(" Rune count:",utf8.RuneCountInString(s))
	bytes := []byte(s)
	fmt.Print(bytes)
	for len(bytes) > 0{
		//解压utf-8编码，并以字节为单位返回符文及其宽度。
		//size是字符的宽度，用来向后移动切片，因为DecodeRune只会拿到第一个切片中的第一个字符。
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	//这样就是一个一个字符的拿到的
	for i, ch := range []rune(s){
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

}