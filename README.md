# Sakana Resume Backend

一个基于 Gin 框架开发的简历访问控制系统，用于控制简历的访问权限。

## 功能特点

- 访问码生成和验证
- JWT 令牌认证
- 访问日志记录
- 自动过期机制
- 一次性使用限制
- 跨域支持
- Docker 容器化部署
- GitHub Actions 自动部署

## 技术栈

- Go 1.20
- Gin Web 框架
- GORM ORM
- MySQL 8.0
- JWT 认证
- Docker & Docker Compose
- GitHub Actions

## 快速开始

### 环境要求

- Docker
- Docker Compose
- Go 1.20+ (本地开发)

### 使用 Docker 部署

1. 克隆项目
```bash
git clone https://github.com/sakana9826/sakana-resume-backend.git
cd sakana-resume-backend
```

2. 配置环境变量
```bash
cp .env.example .env
# 编辑 .env 文件，设置必要的环境变量
```

3. 启动服务
```bash
docker-compose up -d
```

4. 查看服务状态
```bash
docker-compose ps
```

5. 查看日志
```bash
docker-compose logs -f
```

### 本地开发

1. 安装依赖
```bash
go mod download
```

2. 配置环境变量
```bash
cp .env.example .env
# 编辑 .env 文件，设置必要的环境变量
```

3. 运行服务
```bash
go run main.go
```

### 自动部署

项目使用 GitHub Actions 实现自动部署。当向 `main` 分支推送代码时，会自动触发部署流程。

#### 部署流程

1. 连接到服务器
2. 拉取最新代码
3. 停止并删除旧的容器
4. 构建新的镜像
5. 启动新的容器
6. 清理旧的镜像（保留最近5个版本）

#### 配置要求

1. 在 GitHub 仓库的 Settings -> Secrets and variables -> Actions 中添加以下密钥：
   - `SERVER_HOST`：服务器 IP 地址
   - `SERVER_USERNAME`：服务器用户名
   - `SERVER_SSH_KEY`：SSH 私钥

2. 确保服务器上的 `/project/sakana-resume-backend` 目录已存在，并且有正确的权限

3. 确保服务器上已安装 Docker 和 Docker Compose

## API 接口

### 生成访问码

```http
POST /api/generate-access-code
Content-Type: application/json

{
    "expireHours": 24
}
```

响应：
```json
{
    "accessCode": "xxx",
    "expiresAt": "2024-03-22T10:00:00Z"
}
```

### 验证访问码

```http
POST /api/verify-access-code
Content-Type: application/json

{
    "accessCode": "xxx"
}
```

响应：
```json
{
    "token": "jwt-token",
    "expiresAt": "2024-03-22T10:00:00Z"
}
```

## 数据库设计

### access_codes 表
- id: 主键
- code: 访问码（唯一）
- expires_at: 过期时间
- used: 是否已使用
- created_at: 创建时间
- updated_at: 更新时间

### access_logs 表
- id: 主键
- access_code: 使用的访问码
- ip: 访问者IP
- user_agent: 访问者浏览器信息
- accessed_at: 访问时间
- expires_at: 访问权限过期时间
- created_at: 创建时间
- updated_at: 更新时间

## 安全特性

- 访问码使用后立即失效
- JWT 令牌加密存储
- 环境变量配置敏感信息
- 访问日志记录
- 自动过期机制

## 项目结构

```
.
├── config/         # 配置相关
├── controllers/    # 控制器
├── middleware/     # 中间件
├── models/        # 数据模型
├── routes/        # 路由配置
├── utils/         # 工具函数
├── .env           # 环境变量
├── .env.example   # 环境变量示例
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── main.go
```

## 维护命令

- 启动服务：`docker-compose up -d`
- 停止服务：`docker-compose down`
- 重启服务：`docker-compose restart`
- 查看日志：`docker-compose logs -f`
- 重新构建：`docker-compose up -d --build`

## 注意事项

1. 生产环境部署前请修改所有密码和密钥
2. 建议使用 HTTPS 确保通信安全
3. 定期备份数据库
4. 监控服务状态和日志

## 许可证

MIT License 