# Go micro 用户认证权限系统微服务

这是一个基于Go micro + Casbin + Gin + Gorm + JWT 的用户认证和权限的微服务

## 开始

- [目录说明](#目录说明)
- [依赖](#依赖)
- [使用](#使用)

## 生成 proto 文件 (GOPATH 目录里执行)

protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. accbase/srv/auth/proto/auth.proto

## 目录说明

`api` : Restful api 接口实现

`cli` : cli 服务调用

`cmd` : main函数文件目录

`configs`: 项目配置
    
> model.conf 是 Casbin的模型配置（当前模式是RBAC基于角色的访问控制） policy.csv 是 Casbin的策略配置

`go.mod` : 项目依赖管理

`Dockerfile` : Docker 文件

`README.md` : 说明文档

## 依赖
consul 服务发现

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## 使用

运行服务
```
go run cmd/main.go
```

Build a docker image
```
make docker
```

Todo List
1. 用户注册，获取token ❌
2. 用户登录，获取token ❌
3. 根据token，获取用户信息 ❌
4. 根据token，获取用户角色集合 ❌
5. 用户新增角色 ❌
6. 角色新增权限 ❌
7. 角色的增删改查 ❌
8. 权限的增删改查 ❌
9. 系统所有的角色列表 ❌