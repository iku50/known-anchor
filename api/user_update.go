package api

import (
	"fmt"
	"known-anchors/model"
	"known-anchors/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserUpdate(c *gin.Context) {
	service := c.MustGet("service").(*service.ServiceContext)
	uid := fmt.Sprintf("%d", c.MustGet("uid").(uint64))
	if uid != c.Param("userid") {
		c.JSON(http.StatusForbidden, gin.H{
			"data": gin.H{},
			"meta": gin.H{
				"msg":  "权限不足",
				"code": 403,
			},
		})
		return
	}
	userid := c.MustGet("uid").(uint64)
	req := &model.UserUpdateReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": gin.H{},
			"meta": gin.H{
				"msg":  "参数错误",
				"code": 400,
			},
		})
		return
	}
	resp, err := service.UserUpdate(userid, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": gin.H{},
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
			"msg":  "更新成功",
			"code": 200,
		},
	})
}
