package ginuser

import (
	"delivery-food/component"
	"delivery-food/modules/user/usermodel"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		// name, password, email
		var req usermodel.UserCreate
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(200, gin.H{
				"message": "fail",
				"error":   err.Error(),
			})
			return
		}
		fmt.Println(req.Name)
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}
