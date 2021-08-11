package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name string `json:"name"`
}

func GetTags(offset, limit int) ([]Tag, error) {
	var tags []Tag
	var err error
	// 从 offset 开始读取 limit 条
	if limit > 0 && offset > 0 {
		err = db.Find(&tags).Offset(offset).Limit(limit).Error
	} else {
		err = db.Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTag(id int) (*Tag, error) {
	var tag Tag
	if err := db.Where("id = ?", id).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func HasTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	// id 为正数时才表示存在
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func HasTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	// id 为正数时才表示存在
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetTagsCount() (int, error) {
	var count int
	if err := db.Model(&Tag{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func AddTag(name string) error {
	// 根据参数构造 tag 结构体
	tag := Tag{Name: name}
	// 插入记录
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data map[string]interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}
