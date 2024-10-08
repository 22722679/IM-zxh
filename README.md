# IM-zxh
IM即时通讯项目

## 目录
 - [简介](#简介)
 - [项目技术栈](#项目技术栈)
 - [基础环境](#基础环境)
 - [文件目录](#文件目录)
 - [实现功能](#实现功能)
 - [安装步骤](#安装步骤)
 - [项目作者](项目作者)
 - [版本控制](版本控制)
## 简介
  基于gin框架实现的IM即时通讯系统，实现了用户的基本功能(验证码登录/注册/查询/鉴权)，通过MongoDB对用户信息和关系进行存储，实现用户的一对一通讯、群组通讯、消息记录、用户查询/添加/删除好友等多种功能。
## 项目技术栈
 - Go
 - gin
 - MongoDB
 - websocket
 - jwt
## 基础环境
  - Go 1.23.2
  - MongoDB 8.0
  - Docker 26.1.1
## 文件目录
```go
├── define(配置信息)
│   └── define.go
├── helper(Md5码,token,短信验证码)
│   └── helper.go
├── middleware(中间件)
│   └── auth (鉴权)
├── model (结构层)
│   ├── init.go
│   ├── message_basic.go
│   ├── room_basic.go
│   ├── user_basic.go
│   └── user_room.go
├── router(路由)
│   └── router.go
├── service(服务层)
│   ├── chat.go
│   ├── user_basic.go
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
###                 用户模块
####                密码登录
####                发送验证码
####                用户注册
####                用户详情
####                查询用户
## 通讯模块(核心)
####                使用HTTP搭建Websocket服务
####                使用GIN搭建Websockt服务
####                一对一通讯
####                多对多通讯
####                消息列表
####                聊天记录列表

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
5. 在客户端配置相关地址服务端地址即可
## 版本控制
该项目使用Git进行版本管理。您可以在repository参看当前可用版本。


