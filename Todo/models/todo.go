package models

import (
	"fmt"
	"todo/settings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func InitMysql(config *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return
}

func Close() {
	DB.Close()
}

func CreateATodo(todo *Todo) (err error) {
	if err = DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = DB.Save(todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	if err = DB.Where("id = ?", id).Delete(&Todo{}).Error; err != nil {
		return err
	}
	return
}
