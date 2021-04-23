package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//模型绑定和验证
//不同的绑定器要设置不同tag
type Account struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
func Login(c *gin.Context) {
	var account Account
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&account); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK,gin.H{
		"username" : account.Username,
		"password" : account.Password,
	})
}
