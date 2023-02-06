# gin框架小练习
## todo待办清单

### 下载
```bash
git clone https://github.com/mizhexiaoxiao/todo.git
```
### 配置MySQL
1. 在你的数据库中执行以下命令，创建本项目所用的数据库：
```sql
CREATE DATABASE t_todo DEFAULT CHARSET=utf8mb4;
```
2. 在`todo/conf/config.ini`文件中按如下提示配置数据库连接信息。

```ini
port = 9000
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

Mac/Unix：
```bash
./todo conf/config.ini
```
Windows:
```bash
todo.exe conf/config.ini
```

启动之后，使用浏览器打开`http://127.0.0.1:9000/`即可。

### 代码参考
https://github.com/Q1mi/bubble
