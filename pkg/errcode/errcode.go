package errcode

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
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
	OK             = New(200000, "")
	NoSuchResource = New(404000, "No such resource")
	BadRequest     = New(400000, "Bad Request")
	ServerErr      = New(500000, "Server internal error")
	// [600000, ......]
	// 为业务错误码，自行注册，其中前两位为业务编号，后四位为业务错误编号
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
		panic(fmt.Sprintf("errcode: %d already exist", code))
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
	if msg, ok := messages[int(e)]; ok {
		return msg
	}
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
		return ApiError(ServerErr)
	}
	return ApiError(i)
}

// 将其它错误转换为 IApiError（必须是通过 errcode.XXX 引用得到的错误才能被转换）
func Cause(e error) IApiError {
	if e == nil {
		return ApiError(OK)
	}
	ec, ok := errors.Cause(e).(IApiError)
	if ok {
		return ec
	}
	return String(e.Error())
}

func (e ApiError) ToHttpStatusCode() int {
	return e.Code() / 1000
}
