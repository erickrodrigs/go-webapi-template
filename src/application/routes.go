package application

import (
	"github.com/erickrodrigs/go-webapi-template/src/application/controllers"
	"github.com/erickrodrigs/go-webapi-template/src/infrastructure/db"
	"github.com/gin-gonic/gin"
)

// Run ...
func Run(router *gin.Engine) {
	db := db.ConnectDB()
	bookController := controllers.NewBookController(db)

	router.GET("/books", bookController.Index)
	router.POST("/books", bookController.Create)
}
