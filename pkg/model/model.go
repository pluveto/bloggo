package model

/**
 * 全局统一返回结构
 */
type RetWrap struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
