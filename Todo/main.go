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
	// 命令行启动
	if len(os.Args) < 2 {
		fmt.Println("命令使用错误,请使用todo.exe conf/config.ini")
		return
	}
	// 加载配置文件
	err := settings.Init(os.Args[1])
	if err != nil {
		fmt.Printf("加载配置文件失败err: %v", err)
		return
	}
	// 连接数据库，创建模型表
	err = models.InitMysql(settings.Conf.MySQLConfig)
	// err = models.InitMysql((*settings.MySQLConfig)(settings.Conf.PostgreSQLConfig))
	if err != nil {
		panic(err)
	}
	defer models.Close()
	models.DB.AutoMigrate(&models.Todo{})

	// 创建路由
	r := routers.SetupRouter(&f)
	port := fmt.Sprintf("%d", settings.Conf.Port)
	fmt.Printf("http://localhost:%v", settings.Conf.Port)
	r.Run(port)

}
