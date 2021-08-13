package errno

// 自定义的错误码
const (
	Success       = 200
	Error         = 300
	InvalidParams = 400

	IsExisted      = 10001
	GetExistedFail = 10002
	IsNotExist     = 10003
	GetAllFail     = 10004
	CountFail      = 10005
	AddFail        = 10006
	EditFail       = 10007
	DeleteFail     = 10008
)

// 错误码对应的错误消息
var Msg = map[int]string{
	Success:       "成功",
	Error:         "错误",
	InvalidParams: "请求参数错误",

	IsExisted:      "已存在",
	GetExistedFail: "获取失败",
	IsNotExist:     "不存在",
	CountFail:      "统计失败",
	AddFail:        "新增失败",
	EditFail:       "修改失败",
	DeleteFail:     "删除失败",
}
