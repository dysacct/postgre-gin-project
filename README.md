# 快网 CMDB 资产管理系统

高性能、企业级、开箱即用的机器资产管理系统

## 特性
- RESTful API + JWT 认证 + 角色权限
- 多表关联查询 + 全局模糊搜索
- Redis 整页缓存 + 写即失效
- Swagger 在线文档
- Docker 一键部署

## 快速启动
```bash
docker-compose -f docker-compose.prod.yml up -d --build

## 接口文档
http://your-ip:8080/swagger/index.html
```

```bash
步骤 1：安装 swag 工具（只装一次）
go install github.com/swaggo/swag/cmd/swag@latest

# 一定要在项目根目录执行（go.mod 所在目录）
swag init
```