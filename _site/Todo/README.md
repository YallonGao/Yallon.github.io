## 技术栈

前端：vue3

后端：gin+gorm框架

## 无后端启动

从html启动，点击./template/index

## 有后端启动

### 配置MySQL

1.创建数据库

```sql
CREATE DATABASE yallon_todo DEFAULT CHARSET=utf8mb4;
```

2. 在`todo/conf/config.ini`文件中按如下提示配置数据库连接信息。

```ini
port = 9292
release = false

[mysql]
user = 你的数据库用户名
password = 你的数据库密码
host = 你的数据库host地址
port = 你的数据库端口
db = t_todo
```

### 编译
```bash
go build
```

### 执行

Windows:
```bash
todo.exe conf/config.ini
```

启动之后，使用浏览器打开`http://localhost:9292/`。