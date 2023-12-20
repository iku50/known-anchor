package api

import (
	"known-anchors/model"
	"known-anchors/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeckDelete(c *gin.Context) {
	service := c.MustGet("service").(*service.ServiceContext)

	req := &model.DeckDeleteReq{}
	if id, ok := c.Params.Get("id"); ok {
		rId, err := strconv.Atoi(id)
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
		req.Id = uint(rId)
	} else {
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
	_, err := service.DeckDelete(uid, req)
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
		"data": nil,
		"meta": gin.H{
			"msg":  "删除成功",
			"code": 200,
		},
	})
}
