# 基础项目

## 目录结构

输出树形结构: `tree -L 2 -d -I 'vendor'`

```
.
├── bin              #
├── config
│   ├── local
│   └── locales
├── controllers
├── form
├── grpc_service
├── models
├── proto
├── services
│   └── account
└── utils
```

- bin 编译临时存放目录
- config/local 配置文件目录 local|develop|test|demo|prod 对应相对环境
- config/locales 多语言返回信息。通过code对应返回信息
- controllers 访问控制器
- form 请求api参数验证结构体
- grpc_service grpc协议服务层
- models 数据库模型层结构体
- proto grpc定制协议文件和文件go代码文件
- services 服务层。实现业务逻辑层。如：（services/user, services/notice）
- utils 子项目工具

## 数据库

```
create DATABASE gin-base;

CREATE TABLE `t_account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL DEFAULT '' COMMENT '网关uuid',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `actual_name` varchar(255) NOT NULL DEFAULT '' COMMENT '真实名称',
  `created_by_user_Id` int(11) NOT NULL DEFAULT '0' COMMENT '创建id',
  `department` varchar(255) NOT NULL DEFAULT '' COMMENT '部门',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(255) NOT NULL DEFAULT '' COMMENT '手机号',
  `role_id` varchar(255) NOT NULL DEFAULT '' COMMENT '角色id',
  `is_admin` int(1) NOT NULL DEFAULT '0' COMMENT '是否超级管理员(0:否, 1:是)',
  `status` int(1) NOT NULL DEFAULT '1' COMMENT '用户状态(0:禁用, 1:正常)',
  `created_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `valid` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除(0:正常, 1:删除)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_account_security` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL DEFAULT '' COMMENT '网关关联用户',
  `account_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `salt` varchar(255) NOT NULL DEFAULT '' COMMENT '盐',
  `mode` varchar(255) NOT NULL DEFAULT '' COMMENT '加密算法',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `created_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `valid` tinyint(1) NOT NULL DEFAULT '0' COMMENT '逻辑删除(0:正常, 1:删除)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

```
- 修改数据库配置文件: config/local/config.yaml
```
mysql:
  dialect: mysql
  host: 127.0.0.1
  port: 30012
  dbname: gin-base
  user: root
  password: blockshine@123
  charset: utf8
  parseTime: true
  maxIdle: 0
  maxOpen: 200
  debug: false
```
- 改为创建数据库的地址和用户名密码

## 下载并启动

```
mkdir ${path}/src/github.com/hhy5861
export GOPATH=$GOPATH:${path}
cd ${path}/src/github.com/hhy5861
git clone git@github.com:hhy5861/gin-base.git
cd gin-base
go get
sh build.sh
```m

- curl http://127.0.0.1:8080/gin-base/health

## grpc 生成命令
首先要安全grpc protoc工具。 可以先去 https://developers.google.com/protocol-buffers/

```
cd proto
protoc -I ./ ./*.proto --go_out=plugins=grpc:.
```