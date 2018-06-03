package main

import (
	"fmt"
)

//经过测试map比较适合数据量比较小的情况
//还是比较建议使用map比较好理解。
//1、大小开的0xffff，使用数组模拟map比较快，不需要判重，hash等等
//2、放在函数里面每次循环都要从新make
var lastocmap = make([]int, 0xffff)

func maxlenstr(s string) int {
	//代表之前出现过的元素
	//创建map比较消耗时间,修改一下增强性能
	//lastocmap := make(map[rune]int)
	//这样在每次调用函数时，都重新给数组归0了
	//下面这样编译器直接执行了memclr，这样也会花费挺长时间
	for i := range lastocmap {
		lastocmap[i] = 0
	}
	start := 0
	maxlength := 0
	//特别的地方，golang能直接获得[]byte类型的索引和字符的值
	//直接遍历string类型的的到的是按字节来算的index，但是汉字是算三个字节的，顺序就不对了
	//改成rune就是国际版的了
	for i, ch := range []rune(s) {
		if index := lastocmap[ch]; index >= start {
			start = index
		}
		if i - start + 1 > maxlength {
			maxlength = i - start + 1
		}
		lastocmap[ch] = i + 1
	}
	return maxlength
	//字符串里面的函数在strings包里面

}

func main() {
	fmt.Println(maxlenstr(" sdfshdhsdh"))
	fmt.Println(maxlenstr(" sda1561516"))
	fmt.Println(maxlenstr(" ojhnsiofjhni"))
	fmt.Println(maxlenstr(" ojmio*(%^&%*"))
	fmt.Println(maxlenstr(" ijofsiuhfugs"))
	fmt.Println(maxlenstr(" opkoiphj;io"))
	fmt.Println(maxlenstr(" jk kj;i"))
	fmt.Println(maxlenstr(" sda中国文"))
}

