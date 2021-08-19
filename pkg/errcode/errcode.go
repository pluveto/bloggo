package errcode

import (
	"fmt"
	"strconv"
)

var (
	/**
	 *   * 其中 int 是错误码，长度六位，如 404001。
	 *    要求前三位是对应的 HTTP 错误码，后两位递增
	 *   * 其中 string 是错误消息，规定用英文字符串，
	 *     以便将来 i18n 直接 _trans(msg) 得到各语言翻译
	 */
	messages map[int]string = make(map[int]string)
)

var (
	// [000000, 600000) 为公共错误码
	OK                  = New(200000, "")
	ErrBadRequest       = New(400000, "Bad request")
	ErrBadParam         = New(400001, "Bad request: invalid request param, fail to pass validation")
	ErrTokenExpired     = New(400002, "Token has been expired.")
	ErrTokenNotValidYet = New(400003, "Token has been destroyed.")
	ErrTokenMalformed   = New(400004, "Failed to parse token.")
	ErrTokenInvalid     = New(400005, "Token has been expired.")
	ErrWrongPassword    = New(400006, "Incorrect subject or password.")
	ErrTokenMissing     = New(401001, "Missing token.")
	ErrNoSuchResource   = New(404000, "No such resource")
	ErrServerInternal   = New(500000, "Server internal error")
	// [600000, ......]
	// 为业务错误码，自行注册，其中前两位为业务编号，第三位为 HTTP 错误码，后三位为业务错误编号
	//  61  4 0 0 0
	//  ^^  ^
	//  |  客户端错误
	// 业务号 01
)

type ApiError int

/*
 * New 注册一个新的错误码
 */
func New(code int, message string) int {
	if nil == messages {
		// messages = make(map[int]string)
		panic("`messages` map is not initialized!")
	}
	if _, ok := messages[code]; ok {
		panic(fmt.Sprintf("You are registering a duplicated errcode (value is %d), check your code!", code))
	}
	messages[code] = message
	return (code)
}

// IApiError 自定义错误类型
type IApiError interface {
	Error() string
	Code() int
	Message() string
	ToHttpStatusCode() int
}

// 获取错误
func (e ApiError) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

// 获取错误码
func (e ApiError) Code() int { return int(e) }

// 获取错误消息
func (e ApiError) Message() string {
	msg, ok := messages[int(e)]
	if ok {
		return msg
	}
	fmt.Printf("Unregistered errcode: %v\n", int(e))
	return e.Error()
}

// 将字符串错误码转换为 ApiError
func String(e string) ApiError {
	if e == "" {
		return ApiError(OK)
	}
	// try error string
	i, err := strconv.Atoi(e)
	if err != nil {
		fmt.Printf("unknown error code. err: %v\n", err.Error())
		return ApiError(ErrServerInternal)
	}
	return ApiError(i)
}

// 将其它错误转换为 IApiError（必须是通过 errcode.XXX 引用得到的错误才能被转换）
func Cause(e error) ApiError {
	if e == nil {
		return ApiError(OK)
	}
	return String(e.Error())
}

func (e ApiError) ToHttpStatusCode() int {
	// 处理公共错误码
	var c = e.Code() / 1000
	if c < 600 {
		return c
	}
	// 获取业务错误码的状态码
	c = ((c % 10) * 100)
	if c < 600 {
		return c
	}
	// 如果业务错误码的状态码位填写异常，则返回 418
	return 418
}
