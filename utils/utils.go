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
	count := 0
	// c.Query("page") 取回 URL 中的参数，然后再转换成 Int
	// GET /path?id=1234&name=Manu&value=
	// c.Query("id") == "1234"
	// c.Query("name") == "Manu"
	// c.Query("value") == ""
	// c.Query("wtf") == ""
	page, _ := strconv.Atoi(c.Query("page"))
	// page <= 1 时，count 为 0
	if page > 0 {
		// page = 2 时，count = 10
		count = (page - 1) * 10
	}
	return count
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
