package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 跨域允许中间件

func InitCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许 Origin 字段中的域发送请求
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		// 设置预验请求有效期为 86400 秒
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 设置允许请求的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		// 设置允许请求的 Header
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		// 设置拿到除基本字段外的其他字段，如上面的Apitoken, 这里通过引用Access-Control-Expose-Headers，进行配置，效果是一样的。
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Headers")
		// 配置是否可以带认证信息
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// OPTIONS请求返回200
		if c.Request.Method == "OPTIONS" {
			fmt.Println(c.Request.Header)
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
