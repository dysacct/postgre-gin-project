package middleware

import (
	"gin-postgre-project/utils"
	"net/http"
	"strings"

	"ginhub.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "需要登录"})
			c.Abort() // 阻止后续的请求处理
			return
		}

		parts := strings.SplitN(authHeader, " ", 2) // 分割成两部分，第一部分是Bearer，第二部分是Token
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token 格式错误"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token 无效或过期"})
			c.Abort()
			return
		}

		// 把用户信息挂到上下文，后面 handler 可以直接用
		c.Set("username", claims.Username) // 用户名信息也挂到上下文
		c.Set("role", claims.Role)         // 角色信息也挂到上下文
		c.Next()                           // 继续处理后续的请求
	}
}

// gin.HandlerFunc 类型接口是 Gin 框架的中间件接口，用于处理请求的中间件
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "需要管理员权限"})
			c.Abort()
			return
		}

		c.Next()
	}

}
