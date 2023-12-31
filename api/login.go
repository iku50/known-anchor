package api

import (
	"known-anchors/model"
	"known-anchors/service"

	"github.com/gin-gonic/gin"
)

func AuthLoginPost(c *gin.Context) {
	s := c.MustGet("service").(*service.ServiceContext)
	req := model.AuthLoginPostReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"data": gin.H{},
			"meta": gin.H{
				"msg":  "参数错误",
				"code": 400,
			},
		})
		return
	}
	resp, err := s.AuthLoginPost(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"data": gin.H{},
			"meta": gin.H{
				"msg":  err.Error(),
				"code": 500,
			},
		})
		return
	}
	c.JSON(200, gin.H{
		"data": resp,
		"meta": gin.H{
			"msg":  "登录成功",
			"code": 200,
		},
	},
	)
}
