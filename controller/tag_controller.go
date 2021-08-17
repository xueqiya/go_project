package controller

import (
	"github.com/xueqiya/go_project/model"
	"github.com/xueqiya/go_project/utils"
	"github.com/xueqiya/go_project/utils/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	tags, err := model.GetTags(utils.PageNum(c), utils.PageSize)
	if err != nil {
		utils.Response(c, errno.GetAllFail, nil)
		return
	}

	// 计数
	count, err := model.GetTagsCount()
	if err != nil {
		utils.Response(c, errno.CountFail, nil)
		return
	}

	// 填充数据
	data := map[string]interface{}{"lists": tags, "count": count}
	utils.Response(c, errno.Success, data)
}

func GetTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	exist, err := model.HasTagByID(id)
	if err != nil {
		utils.Response(c, errno.GetExistedFail, nil)
		return
	}
	if !exist {
		utils.Response(c, errno.IsNotExist, nil)
		return
	}

	tag, err := model.GetTag(id)
	if err != nil {
		utils.Response(c, errno.GetExistedFail, nil)
		return
	}
	utils.Response(c, errno.Success, tag)
}

type AddTagForm struct {
	Name string `form:"name" valid:"Required;MaxSize(100)"`
}

func AddTag(c *gin.Context) {
	var form AddTagForm
	errCode := utils.BindAndValid(c, &form)
	if errCode != errno.Success {
		utils.Response(c, errCode, nil)
		return
	}

	exist, err := model.HasTagByName(form.Name)
	if err != nil {
		utils.Response(c, errno.GetExistedFail, nil)
		return
	}
	if exist {
		utils.Response(c, errno.IsExisted, nil)
		return
	}

	err = model.AddTag(form.Name)
	if err != nil {
		utils.Response(c, errno.AddFail, nil)
		return
	}

	utils.Response(c, errno.Success, nil)
}

type EditTagForm struct {
	ID   int    `form:"id" valid:"Required;Min(1)"`
	Name string `form:"name" valid:"Required;MaxSize(100)"`
}

func EditTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	form := EditTagForm{ID: id}
	errCode := utils.BindAndValid(c, &form)
	if errCode != errno.Success {
		utils.Response(c, errCode, nil)
		return
	}

	exist, err := model.HasTagByID(form.ID)
	if err != nil {
		utils.Response(c, errno.GetExistedFail, nil)
		return
	}

	if !exist {
		utils.Response(c, errno.IsNotExist, nil)
		return
	}

	data := map[string]interface{}{"name": form.Name}
	err = model.EditTag(form.ID, data)
	if err != nil {
		utils.Response(c, errno.EditFail, nil)
		return
	}

	utils.Response(c, errno.Success, nil)
}

func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	exist, err := model.HasTagByID(id)
	if err != nil {
		utils.Response(c, errno.GetExistedFail, nil)
		return
	}

	if !exist {
		utils.Response(c, errno.IsNotExist, nil)
		return
	}

	if err := model.DeleteTag(id); err != nil {
		utils.Response(c, errno.DeleteFail, nil)
		return
	}

	utils.Response(c, errno.Success, nil)
}
