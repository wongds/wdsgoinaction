package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	result := ""
	//省略初始条件，相当于while
	for ; n > 0; n /= 2 {
		lsb := n % 2//较低位取模结果，下面加在前面
		result = strconv.Itoa(lsb) + result//strconv.Ttoa转换成字符串
	}
	return result
}

//一行一行的来读
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	//省略初始条件和递增条件，不用打；,相当于while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	//条件全都不加，goroutine会用到
	for {
		fmt.Println("死循环")
	}
}


func main() {
	fmt.Println(
		convertToBin(5),//101
		convertToBin(13),//1101,每次对2取模，然后往前加
		convertToBin(72387885),
		convertToBin(0),
	)
	printFile("abc.txt")
	forever()
}