# accbase

> 基于 Casbin 实现访问控制：Golang + Gin + Gorm + Casbin + JWT

##### 项目目标，先熟悉Casbin ，然后将本项目改造成：身份认证+用户权限访问控制微服务

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
| admin  | 管理员表 |
| member  | 会员表 |
| staff  | 员工表 |

#### API 归类
|  功能  | action  | API  | 描述 |
|  ----  | ---- | ----  | ---- |
| 获取角色数组 | GET | /api/roles | 获取所有的角色 |
| 获取主体数组 | GET | /api/subjects | 获取所有的主体（用户或资源）|
| 获取对象数组 | GET | /api/objects | 获取所有的对象（资源或api路径）|
| 获取策略中授权规则 | POST | /api/policy | 获取所有的授权规则（sub can obj）|
| 为用户添加角色 | POST | /api/addRoleForUser | 为用户添加角色 |

