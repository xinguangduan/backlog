package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	DB_NAME     string = "microservices"
	DB_USERNAME string = "root"
	DB_PASSWORD string = "1qaz2wsx"
	DB_HOST     string = "localhost"
	DB_PORT     int    = 3306
	TABLE_BACKLOG  string = "backlog"
)

type (
	BacklogModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed`
	}

	TransformedBacklog struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

var DBUtils *gorm.DB

// 指定表名
func (BacklogModel) TableName() string {
	return TABLE_BACKLOG
}

// 初始化
func init() {
	var err error
	var constr string
	constr = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	DBUtils, err = gorm.Open("mysql", constr)
	if err != nil {
		fmt.Printf("%v",err)
		panic("数据库连接失败")
	}
	DBUtils.AutoMigrate(&BacklogModel{})
}
