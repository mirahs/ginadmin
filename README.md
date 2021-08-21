# ginadmin
golang gin 管理后台模板

## 框架及第三方库
* [gin](https://github.com/gin-gonic/gin)(web框架)
* [pongo2](https://github.com/flosch/pongo2)(模板引擎)
* [pongo2gin](https://gitlab.com/go-box/pongo2gin)(gin框架pongo2模板支持)
* [gorm](https://gorm.io/)(ORM库)
* [ip2region](https://github.com/lionsoul2014/ip2region)(IP地址定位库)
* [layui](https://www.layui.com/)(前端UI框架)

## 配置
* 修改 `main.go` 配置(可以配置 `web 端口`、`MySQL信息`、`后台初始账号密码` 等)
* `MySQL` 创建 `main.go` 配置的数据库, 默认 `ginadmin`

## 启动
```shell
go build ginadmin
./ginadmin
```

## 访问
浏览器输入 `http://localhost:8080/admin/index/login`, 默认初始账号密码 `admin/admin`
