# IM-zxh
IM即时通讯+简化版抖音视频功能项目

## 目录
 - [简介](#简介)
 - [项目技术栈](#项目技术栈)
 - [基础环境](#基础环境)
 - [文件目录](#文件目录)
 - [实现功能](#实现功能)
 - [安装步骤](#安装步骤)
 - [演示视频](#视频流的逻辑演示视频（不包括用户的即时通讯）)
 - [相关技术](#使用到的相关技术)
 - [项目作者](项目作者)
 - [版本控制](版本控制)
## 简介
  基于Gin框架实现的IM即时通讯、视频流系统，实现了用户的基本功能(验证码登录/注册/查询/jwt鉴权)，MongoDB进行用户信息和关系存储， MySQL存储视频和点赞信息，实现用户的视频播放、点赞，评论以及一对一通讯、群组通讯、消息记录保存、用户查询/添加/删除好友等多种功能。
## 项目技术栈
 - Go
 - gin
 - MongoDB
 - websocket
 - zap
 - jwt
## 基础环境
  - Go 1.23.2
  - MongoDB 8.0
  - sqlx v1.3.5
  - Docker 26.1.1
## 文件目录
```go
├── config(配置信息)
│   ├── define.go
│   └── video_user_info.go
├── helper(Md5码,token,短信验证码)
│   └── helper.go
├── middleware(中间件)
│   └── auth (鉴权)
├── model (结构层)
│   ├── favorite.go
│   ├── feed.go
│   ├── init.go
│   ├── message_basic.go
│   ├── room_basic.go
│   ├── user_basic.go
│   └── user_room.go
├── router(路由)
│   └── router.go
├── service(服务层)
│   ├── chat.go
│   ├── favorite.go
│   ├── user_basic.go
│   ├── video_feed.go
│   └── websocket.go
├── test(测试)
│   ├── mongo.go
│   ├── redis.go
│   └── websocket.go
├── go.mod
├── go.sum
├── README.md
└── main.go (主启动文件)
```
## 实现功能
###  用户模块
####    ———— 密码登录
####    ———— 发送验证码
####    ———— 用户注册
####    ———— 用户详情
####    ———— 查询用户
## 通讯模块(核心)
####    ———— 使用HTTP搭建Websocket服务
####    ———— 使用GIN搭建Websockt服务
####    ———— 一对一通讯
####    ———— 多对多通讯
####    ———— 消息列表
####    ———— 聊天记录列表
### 视频流模块
####    ———— 对视频与上述用户进行绑定，显示对应视频的用户信息
####    ———— 非登录游客实现视频观看
####    ———— 实现对用户的视频的点赞功能，并存储点赞相关信息

## 安装步骤
1. 下载源码
```sh
git clone https://github.com/22722679/IM-zxh.git
```
2. 配置相关服务器地址等相关参数(项目为本地IP)
3. 启动服务
```sh
go rum main.gro
```
4. 在客户端配置相关地址服务端地址即可

## 视频流的逻辑演示视频（不包括用户的即时通讯）
[![Watch the video](https://lf3-static.bytednsdoc.com/obj/eden-cn/wthJoabvf_lm_tyvmahsWgpi/ljhwZthlaukjlkulzlp/images/introduce.png)](https://www.douyin.com/video/7274510760062111011)
## 使用到的相关技术
### 框架相关：
 - [Gin](https://gin-gonic.com/docs/)
### 数据库：
 - [MySQL](https://dev.mysql.com/doc/)
 - [MongoDB](https://www.mongodb.com/try/download/shell)
## 版本控制
该项目使用Git进行版本管理。您可以在repository参看当前可用版本。


