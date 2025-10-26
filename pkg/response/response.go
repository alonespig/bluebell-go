package response

import (
	"bluebell/pkg/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	code.Errno
	Data interface{} `json:"data,omitempty"`
}

// json 响应
func JSON(c *gin.Context, httpStatus int, e code.Errno, data interface{}) {
	c.JSON(httpStatus, Response{
		Errno: e,
		Data:  data,
	})
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, code.OK, data)
}
