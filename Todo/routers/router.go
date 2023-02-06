package routers

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"path"
	"todo/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(f *embed.FS) *gin.Engine {
	r := gin.Default()
	t, _ := template.ParseFS(f, "template/*")
	r.SetHTMLTemplate(t)
	// staitc/static/css/app.708ce172.css
	r.StaticFS("/static", http.FS(newPrefixFS("static/", *f)))
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// // 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// // 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}

type PrefixFS struct {
	f      embed.FS
	Prefix string
}

func (fs *PrefixFS) Open(name string) (fs.File, error) {
	return fs.f.Open(path.Join(fs.Prefix, name))
}

//go embed打包后访问静态资源路径变成/static/static
//此处做一下处理去掉第一个/static，让静态资源正常访问
func newPrefixFS(prefix string, f embed.FS) fs.FS {
	return &PrefixFS{
		f:      f,
		Prefix: prefix,
	}
}
