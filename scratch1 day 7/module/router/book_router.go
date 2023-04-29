package router

import (
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/controller"
	"github.com/gin-gonic/gin"
)

func BookRouter(GEngine *gin.Engine, bc controller.BookController) *gin.Engine {
	GEngine.GET("/book", bc.GetAllBook)
	GEngine.GET("/book/:id", bc.GetBookById)
	GEngine.DELETE("/book/:id", bc.DeletedBook)
	GEngine.POST("/book", bc.CreateBook)
	GEngine.PUT("/book/:id", bc.UpdateBook)
	return GEngine
}
