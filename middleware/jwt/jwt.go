package jwt

import (
	"fmt"
	"time"

	"known-anchors/service"
	"known-anchors/util/pwd"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var AuthFaildJSON gin.H

// GetAuthFaildJSON - get auth faild json
func GetAuthFaildJSON() gin.H {
	if AuthFaildJSON == nil {
		AuthFaildJSON = gin.H{
			"data": gin.H{},
			"meta": gin.H{
				"msg":  "权限不足",
				"code": 401,
			},
		}
	}
	return AuthFaildJSON
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uid uint64
		ctx := c.MustGet("service").(*service.ServiceContext).Ctx
		uc := c.MustGet("service").(*service.ServiceContext).DBQuery.User
		tokenStr := c.GetHeader("Authorization")
		// 在 Redis 中检查 token 是否存在
		redisClient := *c.MustGet("service").(*service.ServiceContext).Redis
		uidStr, err := redisClient.Get(ctx, tokenStr)
		if err != nil {
			if err == redis.Nil {
				// Token 不存在于 Redis，进行后续验证
				token, claims, err := pwd.ParseToken(tokenStr)
				if err != nil || !token.Valid {
					c.JSON(401, GetAuthFaildJSON())
					c.AbortWithStatus(401)
					return
				}
				email := claims.Email
				// 检查用户 email 是否存在
				u, err := uc.FindByEmail(email)
				if err != nil {
					c.JSON(401, GetAuthFaildJSON())
					c.AbortWithStatus(401)
					return
				}
				uidStr = fmt.Sprintf("%d", u.ID)
			} else {
				// Redis 查询出错，返回 500
				c.JSON(500, gin.H{
					"data": gin.H{},
					"meta": gin.H{
						"msg":  "Redis 查询出错",
						"code": 500,
					},
				})
				c.AbortWithStatus(500)
				return
			}
		}
		// uidStr 转 uint64
		uid = pwd.StrToUint64(uidStr)
		// 此时 token 验证成功，存储 token 到 Redis
		// 假设 token 过期时间为 expixreTime 秒
		expireTime := time.Duration(3600) * time.Second
		err = redisClient.Set(ctx, tokenStr, uid, expireTime)
		if err != nil {
			c.JSON(500, gin.H{
				"data": gin.H{},
				"meta": gin.H{
					"msg":  "无法存储 Token 到 Redis",
					"code": 500,
				},
			})
			c.AbortWithStatus(500)
			return
		}
		// 将 uid 存储到 context 中
		c.Set("uid", uid)
		// 进行后续操作
		c.Next()
	}
}
