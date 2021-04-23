package user

import (
	"gin-demo/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var userService Services.User
	if err := c.ShouldBind(&userService); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message" : err.Error(),
		})
		return
	}

	id, err  := userService.Insert()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code" : 400,
			"message" : "insert failed",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"code" : 200,
		"message" : "success",
		"data" : id,
	})

}
