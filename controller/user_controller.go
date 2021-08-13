package controller

import (
	"github.com/xueqiya/go_project/model"
	"github.com/xueqiya/go_project/utils"
	"github.com/xueqiya/go_project/utils/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserTo struct {
	*model.User
	Password bool `json:"password,omitempty"`
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

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
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required,max=20,min=6"`
	NikeName string `json:"nike_name" binding:"required,max=20,min=2"`
	Age      int    `json:"age" binding:"number"`
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
	ID       int    `json:"id" binding:"required,min=1"`
	Phone    string `json:"phone" binding:"len=11|len=0"`
	Password string `json:"password" binding:"max=20,min=6|len=0"`
	NikeName string `json:"nike_name" binding:"max=20,min=2|len=0"`
	Age      int    `json:"age" binding:"number"`
	Status   int    `json:"status" binding:"oneof=0 1"`
}

func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	form := EditUserForm{ID: id}
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

	err = model.EditUser(form.ID, form.Phone, form.Password, form.NikeName, form.Age, form.Status)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.EditFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, nil)
}

type LoginForm struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required,max=20,min=6"`
}

func Login(c *gin.Context) {
	var form LoginForm
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

	if !exist {
		utils.Response(c, http.StatusOK, errno.IsNotExist, nil)
		return
	}

	user, err := model.GetUserByPhoneAndPassword(form.Phone, form.Password)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.IsNotExist, user)
		return
	}

	userTo := GetUserTo{User: user}
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, userTo)
}
