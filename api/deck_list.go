package api

import (
	"known-anchors/model"
	"known-anchors/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeckList(c *gin.Context) {
	service := c.MustGet("service").(*service.ServiceContext)

	req := &model.DeckListReq{}
	// 从 params 中获取参数
	if limit := c.Query("limit"); limit != "" {
		var err error
		req.Limit, err = strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"data": nil,
				"meta": gin.H{
					"msg":  "参数错误",
					"code": 400,
				},
			})
			return
		}
	} else {
		req.Limit = 10
	}
	if offset := c.Query("offset"); offset != "" {
		var err error
		req.Offset, err = strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"data": nil,
				"meta": gin.H{
					"msg":  "参数错误",
					"code": 400,
				},
			})
			return
		}
	} else {
		req.Offset = 0
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
