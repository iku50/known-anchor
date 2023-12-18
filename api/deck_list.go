package api

import (
	"known-anchors/model"
	"known-anchors/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeckList(c *gin.Context) {
	service := c.MustGet("service").(*service.ServiceContext)

	req := &model.DeckListReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"meta": gin.H{
				"msg":  "参数错误",
				"code": 400,
			},
		})
		return
	}
	uid := c.MustGet("uid").(uint64)

	resp, err := service.DeckList(uid, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			"meta": gin.H{
				"msg":  err.Error(),
				"code": 500,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
		"meta": gin.H{
			"msg":  "获取成功",
			"code": 200,
		},
	})
}
