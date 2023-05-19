package models

import (
	"fmt"
	"todo/settings"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 初始化数据库
func InitMysql(config *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return
}

// 关闭数据库
func Close() {
	DB.Close()
}

// 创建Todo
func CreateATodo(todo *Todo) (err error) {
	err = DB.Create(&todo).Error
	if err != nil {
		return
	}
	return
}

// 获取全部Todo
func GetAllTodo() (todoList []*Todo, err error) {
	err = DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return
}

// 获取一个Todo
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = DB.Where("id = ?", id).First(todo).Error
	if err != nil {
		return nil, err
	}
	return
}

// 更新一条数据
func UpdateTodo(todo *Todo) (err error) {
	err = DB.Save(todo).Error
	return
}

// 删除一条数据
func DeleteTodo(id string) (err error) {
	err = DB.Where("id = ?", id).Delete(&Todo{}).Error
	if err != nil {
		return err
	}
	return
}
