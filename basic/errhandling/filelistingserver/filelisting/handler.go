package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"
//定义一个用户自己的错误类型
type userError string
//这里不是很明白Error和Message都是返回e的string，有什么区别呢
func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

//用来显示业务逻辑的，应该提出来，提到这里就行了
func HandleFileList(writer http.ResponseWriter, request * http.Request) error{
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start" +  "with " + prefix)
	}
	//默认是list开头的

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//下面这样直接显示内部出错信息给人看，信息看的比较全，但是可能比较底层，给客户看的应该包装一下
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		//因为有包装函数了，因此这里直接有错误扔出去就行了
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	//遇到错误直接扔出去
	if err != nil {
		return err
	}
	//直接打印到屏幕上
	writer.Write(all)
	//没有错误直接返回nil
	return nil
}
