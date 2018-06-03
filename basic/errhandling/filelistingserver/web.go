package main

import (
	"log"
	"net/http"
	"os"

	"wdsgoinaction/basic/errhandling/filelistingserver/filelisting"

)

//为了设置统一的出错处理，appHandeler是一个函数
type appHandler func(writer http.ResponseWriter, request * http.Request) error

//错误处理的包装函数，这样就能隐藏出错的细节，返回想要用户知道的信息
//输入是一个函数，输出也是一个函数
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request){
	//直接返回把错误更改了的原函数
	return func(writer http.ResponseWriter, request *http.Request) {
		//这里直接写一个defer写一个处理的recover函数，使用自己写的recover处理，这样就不会出现太多不友好的报错了，但是要注意打印日志
		//panic错误
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				//log.Print("Panic:", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			//不管怎么样都要log一下,这样就会在终端上打出来，每次出错的日志
			//这个是golang外部包gpmgo里面的，实际上golang自己的库是不带的。golang自己的库应该这么写
			log.Printf("Error occurred " + "handling request: %s", err.Error())
			//log.Warn("Error handling request: %s", err.Error())
			//默认是正确的
			//err.(userError)是golang的类型断言，如果断言正确就会给ok返回true，否则false。如果不加ok直接断言，可能会出现断言失败直接panic
			//处理了usererror
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			//处理了system error
			code := http.StatusOK
			switch  {
			//如果是系统里面不存在的错误
			//已知错误
			case os.IsNotExist(err):
				code = http.StatusNotFound
				//没权限
			case os.IsPermission(err):
				code = http.StatusForbidden
				//不知道的错误
			default:
				code = http.StatusInternalServerError
			}
			//第一个参数是writer，即向谁汇报error，第二个参数可以是err.Error但是这样会将内部的错误暴露出去，因此直接写一句普通的notfound，最后是一个errorcode
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
//实现一个用户类型错误的接口
type userError interface {
	error
	Message() string
}

 func main() {
	//做一个显示文件的webserver
	//这里表示请求一定要是list开头的
	 //http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	 //模拟出错，写错了默认/所有的request都允许handle
	 //事实上http包是有做保护的即实现了recover，但是出错信息还是非常不好看，会报代码的错误，应该自己实现recover
	 http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
