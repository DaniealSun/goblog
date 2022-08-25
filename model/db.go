package model

import (
	"fmt"
	"goblog/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//定义全局的db对象
var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 禁用表名复数映射，新版本写法
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}
	sqlDB, _ := db.DB()
	db.AutoMigrate(&User{}, &Article{}, &Category{})
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
