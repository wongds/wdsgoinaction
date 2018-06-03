package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 1000; i++{
		//一般操作系统并发执行几百个
		//java中要异步io执行才能实现并发执行1000个，但是golang中只需要goroutine就行了，语言原生支持
		//外面的i不安全，所以把i传进去
		//
		go func(ii int) {//race condition//数据访问冲突//如果这里不定义自己的i，有可能出现结果是外面的for结束了，此时i已经超出范围，即10.但是goroutine还没有结束，继续访问i，此时就出现了数据访问冲突。
			for {
				fmt.Printf("Hello form " + "goroutine %d\n", i)
				//a[ii]++
				//goroutine交出控制权，进行切换，一般不会用到，会有其他交出控制权的途径
				runtime.Gosched()
			}
		}(i)
	}
	//事实上，直接a[i]++这样goroutine无法交出控制权，所以main永远执行不到这一句
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
