# mall
### 介绍
mall 是一个基于 golang、 gin、 gorm、 vue3、element plus 开发的商城系统，包括golang后端、 Vue后台管理员前端 ，主要功能有商品管理、订单管理、用户管理。

### 技术选型

前端技术：

| 技术 | 说明 | 相关文档 |
|---|---|---|
| Vue3 | 前端框架 | https://v3.cn.vuejs.org |
| Vue-Router | 页面路由 | https://next.router.vuejs.org |
| Axios | 网络请求库 | https://axios-http.com |
| Vuex | 状态管理 | https://next.vuex.vuejs.org |
| Element Plus | 前端UI组件库 | https://element-plus.org |

后端技术：

| 技术 | 说明 | 相关文档 |
|---|---|---|
| Gin | Web框架 | https://gin-gonic.com |
| Gorm | ORM框架 | https://gorm.io |
| Jwt | 用户认证 | https://github.com/golang-jwt/jwt |
| Captcha | 验证码生成器 | https://github.com/mojocn/base64Captcha |
| Viper | 配置管理 | https://github.com/spf13/viper |

### 运行环境

| 环境 | 版本 |
|---|---|
| go | 1.17.1 |
| mysql | 8.0.16 |
| node.js | 14.13.1 |
| npm | 6.14.8 |

### 安装与启动

**使用git安装：**
```
git clone https://github.com/zchengo/mall.git
```
**相关资源文件：**

数据库文件在 /demo/sql 目录中，推荐使用 Navicat 来运行SQL文件。图片文件在 /demo/image 目录中，请在自己的电脑中新建一个目录，用于存放这些图片。

**修改配置文件：**

配置文件位于 /server/config.yaml，请按实际情况进行修改

**初始化并运行：**

推荐使用Goland打开mall项目，找到Terminal(终端)，执行如下命令。
```
$ cd server
$ go generate
$ go build -o server main.go (windows编译命令为 go build -o server.exe main.go )

# 运行二进制
$ ./server (windows运行命令为 server.exe)

$ cd web
$ npm install
$ npm run serve
```

**使用浏览器访问：** 

成功启动后，即可通过浏览器访问：http://localhost:8080/#/login

用户名: admin 密码: 123

登录后，可以看到商城首页：

![运行结果](https://github.com/zchengo/mall/blob/main/demo/result/demo.png?raw=true)

### 问题反馈

在使用过程中遇到问题，请在评论区反馈！



