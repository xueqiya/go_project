package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/xueqiya/go_project/config"
)

type Model struct {
	ID         int64 `gorm:"primary_key" json:"id"`
	CreatedOn  int   `json:"created_on"`
	ModifiedOn int   `json:"modified_on"`
}

var db *gorm.DB

var dbc = config.Cfg.Database

func Setup() {
	// 构建 DSL
	DSL := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Name, dbc.Charset, dbc.ParseTime, dbc.Loc)

	// 连接到数据库
	var err error
	db, err = gorm.Open(dbc.Dialect, DSL)
	if err != nil {
		log.Fatalf("can't open database err: %v", err)
	}

	db.LogMode(true)

	// 注册回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	// 全局禁用复数表名
	db.SingularTable(true)

	// 最大链接数
	db.DB().SetMaxIdleConns(10)
	// 最大打开链接
	db.DB().SetMaxOpenConns(100)

	// 自动迁移
	db.AutoMigrate(&User{}, &Goods{})
}

func Close() {
	err := db.Close()
	if err != nil {
	}
}

// 注册 gorm 回调函数
// see https://github.com/jinzhu/gorm/blob/master/callback_create.go

// 创建数据时的回调函数
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now().Unix()
		// 如果存在该列
		if createdOn, ok := scope.FieldByName("CreatedOn"); ok {
			// 如果该列的值为空
			if createdOn.IsBlank {
				// 设置该列的值
				if err := createdOn.Set(now); err != nil {
					scope.Log("createdOn.Set() err: %v", err)
				}
			}
		}
		// 如果存在该列
		if modifiedOn, ok := scope.FieldByName("ModifiedOn"); ok {
			// 如果该列的值为空
			if modifiedOn.IsBlank {
				// 设置该列的值
				if err := modifiedOn.Set(now); err != nil {
					scope.Log("modifiedOn.Set() err: %v", err)
				}
			}
		}
	}
}

// 更新数据时的回调函数
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 查找设置有 update_column 标签的列，如果没有
	if _, ok := scope.Get("gorm:update_column"); !ok {
		// 则设置该列的值
		if err := scope.SetColumn("ModifiedOn", time.Now().Unix()); err != nil {
			scope.Log("SetColumn() err: %v", err)
		}
	}
}
