package main

import (
	"embed"
	"fmt"
	"os"
	"todo/models"
	"todo/routers"
	"todo/settings"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//go:embed template/* static/*
var f embed.FS

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./todo conf/config.ini")
		return
	}
	// 加载配置文件
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	//连接数据库
	err := models.InitMysql(settings.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}
	//程序退出时断开数据库连接
	defer models.Close()
	//程序启动自动创建模型表
	models.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter(&f)
	port := fmt.Sprintf(":%d", settings.Conf.Port)
	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
