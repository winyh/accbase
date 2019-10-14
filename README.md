# accbase

> 基于 Casbin 实现访问控制：Golang + Gin + Gorm + Casbin + JWT

##### 项目目标，先熟悉Casbin ，然后将本项目改造成：身份认证+用户权限访问控制微服务

#### 项目下载
克隆项目到 go 的 [工作目录](http://docscn.studygolang.com/doc/code.html) src

#### 代理设置

> 不设置代理很多包拉取不到，原因你懂的

参考阿里云的 [代理设置](http://mirrors.aliyun.com/goproxy/) 


也可以参照这个[设置](https://github.com/goproxy/goproxy.cn) 

#### 数据库设置

DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/casbin?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

root是本地数据库账号， 123456 是你的本地数据库密码

casbin 是数据库名， 项目运行时会自动创建 数据库casbin 和 数据表。数据库名也可以自定义

#### 本地配置文件

在项目根目录新建 .env 文件，这个文件在项目运行时会自动读取相关配置

.env 文件内容如下
```
HTTP_PORT = 8070
```

#### 项目运行

在项目根目录执行 `go run main.go` 会自动下载依赖

`127.0.0.1:8070` 端口就可以访问到项目了

例如：
`127.0.0.1:8070/api/ping`


#### 运行端口修改

可以修改 .env 文件里配置的端口号


#### 功能点
|  功能   | 描述  |
|  ----  | ----  | 
| 角色管理  | 角色的CURD |
| 用户服务  | 用户的CURD |
| 用户认证  | 登录注册授权认证 |

#### 系统预置角色

|  角色   | 描述  | 
|  ----  | ----  |
| admin  | 管理员 |
| member  | 会员 |
| staff  | 员工 |
| finance  | 财务人员 |
| editor  | 编辑 |
| operation  | 运营人员 |

#### 数据表
|  数据表   | 描述  | 
|  ----  | ----  |
| admins  | 管理员表 |
| member  | 会员表 |
| staff  | 员工表 |

> 用户只有在 g 组才被视为用户？

#### 权限API 归类
|  功能  | action  | API  | 描述 |
|  ----  | ---- | ----  | ---- |
| 获取角色数组 | GET | /api/roles | 获取所有的角色 |
| 获取主体数组 | GET | /api/subjects | 获取所有的主体（用户或资源）|
| 获取对象数组 | GET | /api/objects | 获取所有的对象（资源或api路径）|
| 获取策略中授权规则 | POST | /api/policy | 获取所有的授权规则（sub can obj）|
| 为用户添加角色 | POST | /api/addRoleForUser | 为用户添加角色 |
| 为当前策略添加授权规则 | POST | /api/addNamedPolicy | 向当前命名策略添加授权规则|
| 删除用户 | POST | /api/deleteUser | 删除用户，只有在g 组 里才被视为用户，否则也可能是资源|
| 删除角色 | POST | /api/deleteRole | 删除角色 |
| 删除指定用户权限 | POST | /api/deletePermissionForUser | 删除指定用户权限 |



#### Token API 归类
|  功能  | action  | API  | 描述 |
|  ----  | ---- | ----  | ---- |
| 获取token | POST | /api/token | 根据{username,email}获取token |
| 根据token获取用户信息 | GET | /api/user | 获取指定后台用户|


#### 后台用户管理API 归类
|  功能  | action  | API  | 描述 |
|  ----  | ---- | ----  | ---- |
| 新增后台用户 | POST | /api/admin/user | 新增后台用户 |
| 删除后台用户 | DELETE | /api/admin/user/:id | 删除后台用户 |
| 修改后台用户 | PATCH | /api/admin/user/:id | 修改后台用户 |
| 获取所有后台用户 | GET | /api/admin/users | 获取所有后台用户|
| 根据id获取指定后台用户信息 | GET | /api/admin/user/:id | id获取指定后台用户|


#### 后台用户角色API 归类
|  功能  | action  | API  | 描述 |
|  ----  | ---- | ----  | ---- |
| 新增后台用户角色 | POST | /api/admin/role | 新增后台角色 |
| 删除后台用户角色 | DELETE | /api/admin/role/:id | 删除后台角色 |
| 修改后台用户角色 | PATCH | /api/admin/role/:id | 修改后台角色 |
| 获取所有后台用户角色 | GET | /api/admin/roles | 获取所有后台角色 |
| 根据id获取指定后台用户角色信息 | GET | /api/admin/role/:id | id获取指定后台角色 |


> 这个提交主要是为了记录下突破 go micro 了， 顿悟...哈哈哈哈

Api  是对外提供restful. api 服务的 -> 直接调用的是 srv
内部调用 是 客户端去访问 服务端

直接 /rpc 调用的时候会绕过 restful. api（但是 micro api 网关是必须启动的，网关也可配置） -> 直接调用srv (http请求 会自动转成 rpc调用)
