# Swagger API 文档使用指南

## 🎉 Swagger 集成已完成！

你的项目现在已经成功集成了Swagger API文档系统。

## 📋 已完成的配置

### 1. 主要文件修改
- ✅ `main.go`: 添加了Swagger基础注释和路由
- ✅ `handlers/auth.go`: 添加了登录API的Swagger注释
- ✅ `handlers/machine.go`: 添加了机器CRUD操作的Swagger注释
- ✅ `handlers/machine_list.go`: 添加了机器列表API的Swagger注释
- ✅ `docs/`: 自动生成的Swagger文档文件

### 2. API 端点文档
- `POST /api/login` - 用户登录
- `GET /api/machines` - 获取机器列表（支持分页、搜索、过滤）
- `GET /api/machine/{zbx_id}` - 获取单个机器详情
- `POST /api/machine` - 创建机器
- `PUT /api/machine/{zbx_id}` - 更新机器
- `DELETE /api/machine/{zbx_id}` - 删除机器

## 🚀 如何使用

### 1. 启动应用
```bash
cd /Users/dongyasong/Desktop/pyth/fuwuqi/gin-postgre-project

# 确保数据库和Redis正在运行
docker-compose up -d

# 启动应用
go run main.go
```

### 2. 访问Swagger UI
打开浏览器访问：
```
http://localhost:8080/swagger/index.html
```

### 3. API认证
大部分API需要JWT认证：
1. 先调用 `/api/login` 获取token
2. 在Swagger UI中点击右上角的 🔒 按钮
3. 输入: `Bearer YOUR_JWT_TOKEN`
4. 现在可以调用需要认证的API了

## 🔧 开发工作流

### 当你添加新的API时：

1. **添加Swagger注释**
```go
// YourNewAPI godoc
// @Summary      API简述
// @Description  详细描述
// @Tags         标签名
// @Accept       json
// @Produce      json
// @Param        param_name  path/query/body  type  required  "参数描述"
// @Success      200  {object}  ResponseType
// @Security     ApiKeyAuth
// @Router       /your-endpoint [method]
func YourNewAPI(c *gin.Context) {
    // 你的代码
}
```

2. **重新生成文档**
```bash
swag init
```

3. **重启应用**
```bash
go run main.go
```

## 📚 Swagger注释语法参考

### 基本注释
- `@Summary`: API简短描述
- `@Description`: API详细描述
- `@Tags`: API分组标签
- `@Accept`: 接受的内容类型
- `@Produce`: 返回的内容类型

### 参数注释
- `@Param name path string true "描述"` - 路径参数
- `@Param name query string false "描述"` - 查询参数
- `@Param name body Type true "描述"` - 请求体参数

### 响应注释
- `@Success 200 {object} Type` - 成功响应
- `@Failure 400 {object} Type` - 错误响应

### 安全注释
- `@Security ApiKeyAuth` - 需要JWT认证

## 🐛 常见问题

### 1. 编译错误
如果遇到 `LeftDelim` 或 `RightDelim` 错误：
```bash
# 重新生成文档
swag init
# 然后手动删除 docs/docs.go 中的 LeftDelim 和 RightDelim 行
```

### 2. API不显示
确保：
- API函数有正确的Swagger注释
- 运行了 `swag init`
- 重启了应用

### 3. 认证问题
在Swagger UI中：
- 点击右上角的锁图标
- 输入: `Bearer YOUR_JWT_TOKEN`
- 注意 `Bearer ` 前缀和空格

## 📖 更多资源

- [Swaggo官方文档](https://github.com/swaggo/swag)
- [Swagger规范](https://swagger.io/specification/)
- [Gin-Swagger集成](https://github.com/swaggo/gin-swagger)

## ✨ 下一步

你的Swagger集成已经完成！现在你可以：
1. 访问 `http://localhost:8080/swagger/index.html` 查看API文档
2. 使用Swagger UI测试你的API
3. 为新的API添加文档注释
4. 与前端团队分享API文档链接
