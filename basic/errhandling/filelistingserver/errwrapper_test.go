package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

//定义一个用户自己的错误类型
type testinguserError string
//这里不是很明白Error和Message都是返回e的string，有什么区别呢
func (e testinguserError) Error() string {
	return e.Message()
}

func (e testinguserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testinguserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnKnown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
		fmt.Fprintln(writer, "no error")
		return nil
}
//共用handel测试数据
var tests = []struct{
h	appHandler
code int
message string
}{//暂时有问题
{errPanic, 500, "Internal Server Error"},
{errUserError, 400, "user error"},
{errNotFound, 404, "Not Found"},
{errNoPermission, 403, "Forbidden"},
{errUnKnown, 500, "Internal Server Error"},
{noError, 200, "no error"},
}

//测试了一段代码，测试一段函数
//这个是使用假的Request/Response
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://ww.imooc.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

//测试了整个服务器，代码覆盖量笔记哦啊大
//通过启动一个服务器来实现
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string,
	t *testing.T){
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode ||
		body != expectedMsg {
		t.Errorf("expect (%d, %s);" +
			"got (%d, %s)",
			expectedCode, expectedMsg,
			resp.StatusCode, body)
	}
}