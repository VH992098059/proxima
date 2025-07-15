# 灵度英语本 (Proxima English Learning)

一个现代化的英语学习应用，帮助用户高效学习英语单词、句子和听力，并通过AI技术提供个性化学习计划。

## 项目概述

灵比英语本是一个全栈英语学习平台，采用微服务架构，支持移动端、小程序和PC端。项目集成了AI功能，能够根据用户的学习记录生成个性化学习计划，提高学习效率。

## 主要功能

### 核心学习功能
- **单词学习**: 支持单词记忆、拼写练习、词汇测试
- **句子学习**: 提供句子理解、语法分析、翻译练习
- **听力学习**: 音频播放、听写练习、发音对比
- **学习进度追踪**: 详细记录学习历史和掌握程度
- **错词本**: 自动收集错误单词，便于复习巩固

### AI助手功能
- **智能学习计划**: 基于学习记录生成个性化计划
- **学习建议**: AI分析学习薄弱环节，提供针对性建议
- **进度预测**: 预估学习目标完成时间

### 用户系统
- **用户认证**: 注册、登录、权限管理
- **个人中心**: 学习统计、成就系统、设置管理
- **多端同步**: 支持移动端、小程序、PC端数据同步

## 技术架构

### 后端技术栈
- **框架**: GoFrame (Go语言企业级框架)
- **架构**: 微服务架构
- **数据库**: MongoDB (文档数据库)
- **缓存**: Redis (内存数据库)
- **服务发现**: Etcd
- **通信协议**: gRPC
- **容器化**: Docker + Kubernetes

### 前端技术栈
- **构建工具**: Vite
- **框架**: React + TypeScript
- **路由**: React Router DOM
- **UI组件**: Ant Design
- **包管理**: pnpm
- **状态管理**: React Hooks

### AI集成
- **AI服务**: OpenAI API
- **技术**: RAG (检索增强生成)
- **实时通信**: SSE (Server-Sent Events)

## 项目结构

```
proxima/
├── app/                    # 后端微服务
│   ├── gateway/           # API网关服务
│   ├── user/              # 用户服务
│   ├── word/              # 单词服务
│   └── einoPackage/       # AI服务包
├── front/                 # 前端应用
│   ├── src/
│   │   ├── components/    # 组件库
│   │   ├── pages/         # 页面组件
│   │   ├── api/           # API接口
│   │   ├── hooks/         # 自定义Hooks
│   │   └── router/        # 路由配置
│   └── package.json
├── utility/               # 公共工具库
├── hack/                  # 构建脚本
└── README.MD
```

## 快速开始

### 环境要求
- Go 1.23.1+
- Node.js 18+
- pnpm 8+
- MongoDB 5.0+
- Redis 6.0+
- Etcd 3.5+
- Docker & Docker Compose (推荐)

### 方式一：Docker Compose 启动（推荐）
```bash
# 启动所有服务（包括基础设施）
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f

# 停止所有服务
docker-compose down
```

### 方式二：本地开发启动

#### 启动基础设施
```bash
# 启动 MongoDB, Redis, Etcd
docker-compose up -d mongodb redis etcd
```

#### 启动后端服务
```bash
# 启动网关服务
cd app/gateway
go run main.go

# 启动用户服务
cd app/user
go run main.go

# 启动单词服务
cd app/word
go run main.go
```

#### 启动前端
```bash
cd front
pnpm install
pnpm dev
```

### 访问地址
- 前端应用: http://localhost:5173
- API网关: http://localhost:8000
- MongoDB: localhost:27017
- Redis: localhost:6379
- Etcd: localhost:2379

## 开发进度

### 已完成功能
- ✅ 用户认证系统（注册、登录、JWT验证）
- ✅ 微服务架构搭建（Gateway、User、Word服务）
- ✅ 前端基础框架（React + TypeScript + Ant Design）
- ✅ API接口对接和类型安全
- ✅ Docker容器化部署
- ✅ 服务发现和负载均衡
- ✅ 统一日志和错误处理

### 开发中功能
- 🔄 单词学习核心功能
- 🔄 AI学习计划生成
- 🔄 学习进度追踪
- 🔄 错词本功能

### 计划功能
- 📋 听力练习模块
- 📋 移动端适配
- 📋 小程序版本
- 📋 高级AI功能
- 📋 社交学习功能
- 📋 离线学习支持

## 特色功能

- **智能化**: AI驱动的个性化学习体验
- **多端支持**: 一套代码，多端运行
- **微服务**: 高可用、可扩展的架构设计
- **现代化**: 采用最新的技术栈和开发理念
- **用户友好**: 简洁直观的用户界面

## 部署说明

### 生产环境部署
```bash
# 使用生产配置启动
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

### 环境变量配置
项目支持通过环境变量进行配置：
- `.env`: 本地开发环境配置
- `.env.docker`: Docker环境配置

主要配置项：
- `Model_Type`: AI模型类型
- `Openai_API_Key`: OpenAI API密钥
- `Base_url`: AI服务基础URL
- `ETCD_ENDPOINTS`: Etcd服务地址

## 测试

```bash
# 运行后端测试
cd app/user
go test -v

# 运行前端测试
cd front
pnpm test
```

## 贡献指南

欢迎提交Issue和Pull Request来帮助改进项目。

### 开发规范
- 遵循Go和TypeScript代码规范
- 提交前请运行测试
- 提交信息请使用中文描述

## 许可证

MIT License

---

**灵比英语本** - 让英语学习更智能、更高效！
