package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Phone    string `json:"phone"`
	Password string `json:"password"`
	NikeName string `json:"nike_name"`
	Age      int    `json:"age"`
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
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddUser(phone, password, nikeName string, age int) error {
	user := User{Phone: phone, Password: password, NikeName: nikeName, Age: age, Status: 1}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func EditUser(id int, phone, password, nikeName string, age, status int) error {
	user := User{Phone: phone, Password: password, NikeName: nikeName, Age: age, Status: status}
	if err := db.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByPhoneAndPassword(phone, password string) (*User, error) {
	var user User
	if err := db.Where("phone = ? And password = ?", phone, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
