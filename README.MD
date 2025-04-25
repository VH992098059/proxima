# 灵比英语本（pro版本）

对Goframe社区教程对功能实现进一步扩展，主要包括账号服务与单词服务的功能完善。后期会加入rabbitMQ、es搜索引擎和代码优化的处理，首要任务将账号服务和单词服务的基本功能

## 项目结构



<p style="text-align:center">类图</p>

```
│  main.go
│  Makefile
│  README.MD
│
├─api
│  ├─account
│  │  └─v1
│  │          account.pb.go
│  │          account_grpc.pb.go
│  │
│  ├─account_edit_delete
│  │  └─v1
│  │          account_edit_delete.pb.go
│  │          account_edit_delete_grpc.pb.go
│  │
│  ├─account_logout
│  │  └─v1
│  │          account_logout.pb.go
│  │          account_logout_grpc.pb.go
│  │
│  ├─pbentity
│  │      users.pb.go
│  │
│  └─user_msg
│      └─v1
│              user_msg.pb.go
│              user_msg_grpc.pb.go
│
├─hack
│      config.yaml
│      hack-cli.mk
│      hack.mk
│
├─internal
│  ├─cmd
│  │      cmd.go
│  │
│  ├─consts
│  │      consts.go
│  │
│  ├─controller
│  │  ├─account
│  │  │      account.go
│  │  │
│  │  ├─account_edit_delete
│  │  │      account_edit_delete.go
│  │  │
│  │  ├─account_logout
│  │  │      account_logout.go
│  │  │
│  │  └─user_msg
│  │          user_msg.go
│  │
│  ├─dao
│  │  │  users.go
│  │  │
│  │  └─internal
│  │          users.go
│  │
│  ├─logic
│  │  │  .gitkeep
│  │  │
│  │  ├─account
│  │  │      account.go
│  │  │
│  │  ├─account_edit_delete
│  │  │      account_edit_delete.go
│  │  │
│  │  ├─account_logout
│  │  │      account_logout.go
│  │  │
│  │  └─user_msg
│  │          user_msg.go
│  │
│  ├─model
│  │  │
│  │  ├─do
│  │  │      .gitkeep
│  │  │      users.go
│  │  │
│  │  └─entity
│  │          .gitkeep
│  │          users.go
│  │
│  ├─packed
│  │      packed.go
│  │
│  └─service
│
├─manifest
│  ├─config
│  │      config.yaml
│  │      etcd.yaml
│  │
│  └─protobuf
│      │
│      ├─account
│      │  └─v1
│      │          account.proto
│      │
│      ├─account_edit_delete
│      │  └─v1
│      │          account_edit_delete.proto
│      │
│      ├─account_logout
│      │  └─v1
│      │          account_logout.proto
│      │
│      ├─pbentity
│      │      users.proto
│      │
│      └─user_msg
│          └─v1
│                  user_msg.proto
│
│
└─utility
        passwordEncrypt.go


```



<p style="text-align:center">用户模块树图</p>