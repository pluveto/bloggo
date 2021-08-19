package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pluveto/bloggo/pkg/bloggo/service"
	"github.com/pluveto/bloggo/pkg/errcode"
	"github.com/pluveto/bloggo/pkg/model"
)

func abort(c *gin.Context, apiErr errcode.ApiError) {
	var resp = &model.RetWrap{
		Code:    apiErr.Code(),
		Message: apiErr.Message(),
		Data:    nil,
	}
	c.JSON(apiErr.ToHttpStatusCode(), resp)
	c.Abort()
}

func JWTAuth(service *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		authHeader := c.Request.Header.Get("authorization")
		authHeaderTrimed := strings.TrimSpace(authHeader)
		if len(authHeaderTrimed) == 0 {
			abort(c, errcode.ApiError(errcode.ErrTokenMissing))
			return
		}
		if !strings.HasPrefix(authHeaderTrimed, "Bearer") && !strings.HasPrefix(authHeaderTrimed, "bearer") {
			abort(c, errcode.ApiError(errcode.ErrTokenInvalid))
			return
		}
		token := authHeaderTrimed[len("bearer"):]
		if token == "" {
			abort(c, errcode.ApiError(errcode.ErrTokenMissing))
			return
		}
		token = strings.TrimSpace(token)
		if service.TokenIsDestroyed(token) {
			abort(c, errcode.ApiError(errcode.ErrTokenInvalid))
			return
		}
		// parseToken 解析token包含的信息
		claims, err := service.TokenParse(token)
		if err != nil {
			abort(c, errcode.Cause(err))
			return
		}
		c.Set("claims", claims)
		id, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			abort(c, errcode.ApiError(errcode.ErrServerInternal))
			fmt.Printf("Invalid claims!!!\n")
			return
		}
		c.Set("uid", id)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("TODO: KEYHERE"),
	}
}
