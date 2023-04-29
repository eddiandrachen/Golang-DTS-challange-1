package server

import (
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/controller"
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/repository/book"
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/router"
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/service"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() {
	ginServer := gin.Default()

	br := book.NewBookRepository()
	bs := service.NewBookService(br)
	bc := controller.NewBookController(bs)

	ginServer = router.BookRouter(ginServer, bc)

	ginServer.Run(":30102")
}
