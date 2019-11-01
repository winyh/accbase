# Go micro 用户认证权限系统微服务

这是一个基于Go micro + Casbin + Gin + Gorm + JWT 的用户认证和权限的微服务

## 开始

- [目录说明](#目录说明)
- [依赖](#依赖)
- [Usage](#usage)


protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. proto/test/test.proto

## 目录说明

`api` : Restful api 接口实现

`cmd` : main函数文件目录

`configs`: 项目配置
    
    model.conf 是 Casbin的模型配置（当前模式是RBAC基于角色的访问控制）
    policy 是 Casbin的策略配置

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