package controller

import (
	"github.com/xueqiya/go_project/model"
	"github.com/xueqiya/go_project/utils"
	"github.com/xueqiya/go_project/utils/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllGoods(c *gin.Context) {
	goods, err := model.GetAllGoods(utils.PageNum(c), utils.PageSize)
	if err != nil {
		utils.Response(c, errno.GetAllFail, nil)
		return
	}
	utils.Response(c, errno.Success, goods)
}

func GetGoods(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	exist, err := model.HasGoodsByID(id)
	if err != nil {
		utils.Response(c, errno.GetExistedFail, nil)
		return
	}
	if !exist {
		utils.Response(c, errno.IsNotExist, nil)
		return
	}

	goods, err := model.GetGoods(id)
	if err != nil {
		utils.Response(c, errno.GetExistedFail, nil)
		return
	}
	utils.Response(c, errno.Success, goods)
}

type AddGoodsForm struct {
	UserId   int    `json:"user_id" binding:"required,max=20,min=1"`
	Price    int    `json:"price" binding:"required,number"`
	Keyword  string `json:"keyword" binding:"required,max=20,min=2"`
	Content  string `json:"content" binding:"required,max=200,min=2"`
	Location string `json:"location" binding:"required,max=100,min=2"`
	Address  string `json:"address" binding:"required,max=100,min=2"`
}

func AddGoods(c *gin.Context) {
	var form AddGoodsForm
	errCode := utils.BindAndValid(c, &form)
	if errCode != errno.Success {
		utils.Response(c, errCode, nil)
		return
	}

	exist, err := model.HasUserByID(form.UserId)
	if err != nil {
		utils.Response(c, errno.GetExistedFail, nil)
		return
	}
	if !exist {
		utils.Response(c, errno.IsNotExist, nil)
		return
	}

	err = model.AddGoods(form.UserId, form.Price, form.Keyword, form.Content, form.Location, form.Address)
	if err != nil {
		utils.Response(c, errno.AddFail, nil)
		return
	}

	utils.Response(c, errno.Success, nil)
}

//type EditGoodsForm struct {
//	ID   int    `form:"id" valid:"Required;Min(1)"`
//	Name string `form:"name" valid:"Required;MaxSize(100)"`
//}
//
//func EditGoods(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//	form := EditGoodsForm{ID: id}
//	errCode := utils.BindAndValid(c, &form)
//	if errCode != errno.Success {
//		utils.Response(c, errCode, nil)
//		return
//	}
//
//	exist, err := model.HasGoodsByID(form.ID)
//	if err != nil {
//		utils.Response(c, errno.GetExistedFail, nil)
//		return
//	}
//
//	if !exist {
//		utils.Response(c, errno.IsNotExist, nil)
//		return
//	}
//
//	data := map[string]interface{}{"name": form.Name}
//	err = model.EditGoods(form.ID, data)
//	if err != nil {
//		utils.Response(c, errno.EditFail, nil)
//		return
//	}
//
//	utils.Response(c, errno.Success, nil)
//}
//
//func DeleteGoods(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	exist, err := model.HasGoodsByID(id)
//	if err != nil {
//		utils.Response(c, errno.GetExistedFail, nil)
//		return
//	}
//
//	if !exist {
//		utils.Response(c, errno.IsNotExist, nil)
//		return
//	}
//
//	if err := model.DeleteGoods(id); err != nil {
//		utils.Response(c, errno.DeleteFail, nil)
//		return
//	}
//
//	utils.Response(c, errno.Success, nil)
//}
