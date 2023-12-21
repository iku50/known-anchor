package api

import (
	"known-anchors/model"
	"known-anchors/service"

	"known-anchors/dal/mq"

	"github.com/gin-gonic/gin"
)

func AuthRegisterPost(c *gin.Context) {
	s := c.MustGet("service").(*service.ServiceContext)
	m := c.MustGet("mailproducer").(chan mq.Message)
	req := model.AuthRegisterPostReq{}
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
	resp, err := s.AuthRegisterPost(c.Request.Context(), &req, m)
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
			"msg":  "注册成功",
			"code": 200,
		},
	},
	)
}
