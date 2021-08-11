package controller

import (
	"github.com/linehk/gin-blog/model"
	"github.com/linehk/gin-blog/utils"
	"github.com/linehk/gin-blog/utils/errno"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context) {
	tags, err := model.GetTags(utils.PageNum(c), utils.PageSize)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetAllTagFail, nil)
		return
	}

	// 计数
	count, err := model.GetTagsCount()
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.CountTagFail, nil)
		return
	}

	// 填充数据
	data := map[string]interface{}{"lists": tags, "count": count}
	utils.Response(c, http.StatusOK, errno.Success, data)
}

func GetTag(c *gin.Context) {
	// 获取 id
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	// 表单验证错误
	if valid.HasErrors() {
		utils.LogErrors(valid.Errors)
		utils.Response(c, http.StatusBadRequest, errno.InvalidParams, nil)
	}

	exist, err := model.HasTagByID(id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedTagFail, nil)
		return
	}
	if !exist {
		utils.Response(c, http.StatusOK, errno.TagIsNotExist, nil)
		return
	}

	tag, err := model.GetTag(id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedTagFail, nil)
		return
	}
	utils.Response(c, http.StatusOK, errno.Success, tag)
}

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
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
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedTagFail, nil)
		return
	}
	if exist {
		utils.Response(c, http.StatusOK, errno.TagNameIsExisted, nil)
		return
	}

	err = model.AddTag(form.Name, form.CreatedBy)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.AddTagFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, nil)
}

type EditTagForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

func EditTag(c *gin.Context) {
	form := EditTagForm{ID: com.StrTo(c.Param("id")).MustInt()}
	httpCode, errCode := utils.BindAndValid(c, &form)
	if errCode != errno.Success {
		utils.Response(c, httpCode, errCode, nil)
		return
	}

	exist, err := model.HasTagByID(form.ID)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedTagFail, nil)
		return
	}

	if !exist {
		utils.Response(c, http.StatusOK, errno.TagIsNotExist, nil)
		return
	}

	data := map[string]interface{}{
		"modified_by": form.ModifiedBy,
		"name":        form.Name,
	}
	err = model.EditTag(form.ID, data)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.EditTagFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, nil)
}

func DeleteTag(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		utils.LogErrors(valid.Errors)
		utils.Response(c, http.StatusBadRequest, errno.InvalidParams, nil)
	}

	exist, err := model.HasTagByID(id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.GetExistedTagFail, nil)
		return
	}

	if !exist {
		utils.Response(c, http.StatusOK, errno.TagIsNotExist, nil)
		return
	}

	if err := model.DeleteTag(id); err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.DeleteTagFail, nil)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, nil)
}
