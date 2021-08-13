package controller

import (
	"github.com/xueqiya/go_project/model"
	"github.com/xueqiya/go_project/utils"
	"github.com/xueqiya/go_project/utils/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	tags, err := model.GetTags(utils.PageNum(c), utils.PageSize)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetAllFail, nil)
		return
	}

	// 计数
	count, err := model.GetTagsCount()
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.CountFail, nil)
		return
	}

	// 填充数据
	data := map[string]interface{}{"lists": tags, "count": count}
	utils.Response(c, http.StatusOK, errno.Success, data)
}

func GetTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	exist, err := model.HasTagByID(id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}
	if !exist {
		utils.Response(c, http.StatusOK, errno.IsNotExist, nil)
		return
	}

	tag, err := model.GetTag(id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}
	utils.Response(c, http.StatusOK, errno.Success, tag)
}

type AddTagForm struct {
	Name string `form:"name" valid:"Required;MaxSize(100)"`
}

func AddTag(c *gin.Context) {
	var form AddTagForm
	httpCode, errCode := utils.BindAndValid(c, &form)
	if errCode != errno.Success {
		utils.Response(c, httpCode, errCode, nil)
		return
	}

	exist, err := model.HasTagByName(form.Name)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}
	if exist {
		utils.Response(c, http.StatusOK, errno.IsExisted, nil)
		return
	}

	err = model.AddTag(form.Name)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.AddFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, nil)
}

type EditTagForm struct {
	ID   int    `form:"id" valid:"Required;Min(1)"`
	Name string `form:"name" valid:"Required;MaxSize(100)"`
}

func EditTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	form := EditTagForm{ID: id}
	httpCode, errCode := utils.BindAndValid(c, &form)
	if errCode != errno.Success {
		utils.Response(c, httpCode, errCode, nil)
		return
	}

	exist, err := model.HasTagByID(form.ID)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}

	if !exist {
		utils.Response(c, http.StatusOK, errno.IsNotExist, nil)
		return
	}

	data := map[string]interface{}{"name": form.Name}
	err = model.EditTag(form.ID, data)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.EditFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, nil)
}

func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	exist, err := model.HasTagByID(id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedFail, nil)
		return
	}

	if !exist {
		utils.Response(c, http.StatusOK, errno.IsNotExist, nil)
		return
	}

	if err := model.DeleteTag(id); err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.DeleteFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, nil)
}
