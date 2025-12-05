# 环境配置说明

## 创建 .env 文件

请在项目根目录创建 `.env` 文件，内容如下：

```bash
# .env.production
GIN_MODE=release
DB_HOST=3417.s.kuaicdn.cn
DB_PORT=34179
DB_USER=kuaicdn
DB_PASSWORD=001002
DB_NAME=machine_info
DB_SSLMODE=disable
JWT_SECRET=23fb5304cf431ec5fe178a7e3b144c5c1442620c397a0f8753a1dc63e817f38f
JWT_EXPIRE_HOURS=720
```

## JWT_SECRET 说明

### 什么是 JWT_SECRET？
- JWT_SECRET 是用于签名和验证 JWT 令牌的密钥
- 它确保令牌的完整性和安全性
- **必须保密**，不能泄露给客户端

### 如何生成新的 JWT_SECRET？

#### 方法1：使用 OpenSSL
```bash
openssl rand -hex 32
```

#### 方法2：使用项目中的生成器
```bash
go run generate_jwt_secret.go
```

#### 方法3：使用 Python
```bash
python3 -c "import secrets; print(secrets.token_hex(32))"
```

### 安全建议

1. **定期更换**：建议每3-6个月更换一次JWT密钥
2. **环境隔离**：生产环境和开发环境使用不同的密钥
3. **长度要求**：至少32字节（64个十六进制字符）
4. **随机性**：使用加密安全的随机数生成器

### 配置项说明

- `GIN_MODE`: Gin框架运行模式（debug/release）
- `DB_HOST`: 数据库主机地址
- `DB_PORT`: 数据库端口
- `DB_USER`: 数据库用户名
- `DB_PASSWORD`: 数据库密码
- `DB_NAME`: 数据库名称
- `DB_SSLMODE`: SSL模式（disable/require）
- `JWT_SECRET`: JWT签名密钥（64字符十六进制字符串）
- `JWT_EXPIRE_HOURS`: JWT令牌过期时间（小时）
