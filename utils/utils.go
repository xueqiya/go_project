package utils

import (
	"github.com/xueqiya/go_project/utils/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PageSize 每页 Article 的数量
var PageSize = 10

func PageNum(c *gin.Context) int {
	offset := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		offset = (page - 1) * PageSize
	}
	return offset
}

// Resp 统一返回格式
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response 根据数据返回响应
func Response(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(http.StatusOK, Resp{Code: httpCode, Msg: errno.Msg[httpCode], Data: data})
}

// BindAndValid 绑定并验证表单
func BindAndValid(c *gin.Context, form interface{}) int {
	// c.Bind(form) 会根据 Content-Type 选择 binding
	err := c.Bind(form)
	if err != nil {
		return errno.InvalidParams
	} else {
		return errno.Success
	}
}
