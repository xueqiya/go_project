package model

import (
	"github.com/jinzhu/gorm"
)

type Goods struct {
	Model
	FkUserGoods int    `json:"fk_user_goods"`
	Price       int    `json:"price"`
	Keyword     string `json:"keyword"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	Status      int    `json:"status"`
}

type GoodsTo struct {
	Goods
	NikeName string `json:"nike_name"`
}

func GetAllGoods(offset, limit int) ([]GoodsTo, error) {
	var goods []GoodsTo
	var err = db.Table("goods").Select("user.nike_name,goods.*").Joins("join user on user.id = goods.fk_user_goods").Order("goods.created_on desc").Offset(offset).Limit(limit).Find(&goods).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return goods, nil
}

func GetGoods(id int) (*Goods, error) {
	var goods Goods
	if err := db.Where("id = ?", id).First(&goods).Error; err != nil {
		return nil, err
	}
	return &goods, nil
}

func HasGoodsByID(id int) (bool, error) {
	var goods Goods
	err := db.Select("id").Where("id = ?", id).First(&goods).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	// id 为正数时才表示存在
	if goods.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddGoods(userId, price int, keyword, content, location, address string) error {
	// 根据参数构造 goods 结构体
	goods := Goods{FkUserGoods: userId, Price: price, Keyword: keyword, Content: content, Location: location, Address: address, Status: 1}
	// 插入记录
	if err := db.Create(&goods).Error; err != nil {
		return err
	}
	return nil
}

//func EditGoods(id int, data map[string]interface{}) error {
//	if err := db.Model(&Goods{}).Where("id = ?", id).Updates(data).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func DeleteGoods(id int) error {
//	if err := db.Where("id = ?", id).Delete(&Goods{}).Error; err != nil {
//		return err
//	}
//	return nil
//}
