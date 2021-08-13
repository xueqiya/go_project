package controller

import (
	"github.com/xueqiya/go_project/model"
	"github.com/xueqiya/go_project/utils"
	"github.com/xueqiya/go_project/utils/errno"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type GetUserTo struct {
	*model.User
	Password bool `json:"password,omitempty"`
}

func GetUser(c *gin.Context) {
	// 获取 id
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	// 表单验证错误
	if valid.HasErrors() {
		utils.LogErrors(valid.Errors)
		utils.Response(c, http.StatusBadRequest, errno.InvalidParams, nil)
	}

	exist, err := model.HasUserByID(id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}
	if !exist {
		utils.Response(c, http.StatusOK, errno.IsNotExist, nil)
		return
	}

	user, err := model.GetUser(id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}

	userTo := GetUserTo{User: user}
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, userTo)
}

type AddUserForm struct {
	Phone    string `json:"phone" valid:"Required;Phone;Length(11)"`
	Password string `json:"password" valid:"Required;MaxSize(20);MinSize(6)"`
	NikeName string `json:"nike_name" valid:"Required;MaxSize(20),MinSize(2)"`
	Age      string `json:"age" valid:"Numeric"`
}

func AddUser(c *gin.Context) {
	var form AddUserForm
	httpCode, errCode := utils.BindAndValid(c, &form)
	if errCode != errno.Success {
		utils.Response(c, httpCode, errCode, nil)
		return
	}

	exist, err := model.HasUserByPhone(form.Phone)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}
	if exist {
		utils.Response(c, http.StatusOK, errno.IsExisted, nil)
		return
	}

	err = model.AddUser(form.Phone, form.Password, form.NikeName, form.Age)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.AddFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, nil)
}

type EditUserForm struct {
	ID       int    `form:"id" valid:"Required;Min(1)"`
	Phone    string `json:"phone" valid:"Phone;Length(11)"`
	Password string `json:"password" valid:"MaxSize(20);MinSize(6)"`
	NikeName string `json:"nike_name" valid:"MaxSize(20),MinSize(2)"`
	Age      string `json:"age" valid:"Numeric"`
	Status   string `json:"status" valid:"Numeric"`
}

func EditUser(c *gin.Context) {
	form := EditUserForm{ID: com.StrTo(c.Param("id")).MustInt()}
	httpCode, errCode := utils.BindAndValid(c, &form)
	if errCode != errno.Success {
		utils.Response(c, httpCode, errCode, nil)
		return
	}

	exist, err := model.HasUserByID(form.ID)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}

	if !exist {
		utils.Response(c, http.StatusOK, errno.IsNotExist, nil)
		return
	}

	err = model.EditUser(form.ID, form.Phone, form.Password, form.NikeName, form.Age)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.EditFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, nil)
}
