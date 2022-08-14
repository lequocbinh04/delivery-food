package main

import (
	"delivery-food/component"
	"delivery-food/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "root:EXdgAahbXpqjj4M71q1VkizjKIYDBnQm5A39kFiTKOhuKeT37XxkNBgN12TMX7SD@tcp(206.189.41.189:3306)/delivery-food?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	appCtx := component.NewAppContext(db)
	if err := runService(appCtx); err != nil {
		log.Fatalln(err)
	}
}

func runService(appCtx component.AppContext) error {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	authRoute := v1.Group("/auth")
	{
		authRoute.POST("/register", ginuser.Register(appCtx))
	}

	return r.Run()
}
