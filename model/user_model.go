package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Phone    string `json:"phone"`
	Password string `json:"password"`
	NikeName string `json:"nike_name"`
	Age      string `json:"age"`
	Status   int    `json:"status"`
}

func GetUser(id int) (*User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func HasUserByID(id int) (bool, error) {
	var user User
	err := db.Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	// id 为正数时才表示存在
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func HasUserByPhone(phone string) (bool, error) {
	var user User
	err := db.Select("id").Where("phone = ?", phone).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	// id 为正数时才表示存在
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddUser(phone, password, nikeName, age string) error {
	// 根据参数构造 user 结构体
	user := User{Phone: phone, Password: password, NikeName: nikeName, Age: age, Status: 1}

	// 插入记录
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func EditUser(id int, phone, password, nikeName, age string) error {
	user := User{Phone: phone, Password: password, NikeName: nikeName, Age: age, Status: 1}
	if err := db.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
