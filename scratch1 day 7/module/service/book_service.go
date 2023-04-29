package service

import (
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/model"
	"github.com/gin-gonic/gin"
)

type BookService interface {
	Insert(ctx *gin.Context, request model.BookCreateRequest) model.BookResponse
	Update(ctx *gin.Context, request model.BookUpdateRequest) (model.BookResponse, error)
	Delete(ctx *gin.Context, id int) error
	GetAll(ctx *gin.Context) []model.BookResponse
	GetById(ctx *gin.Context, id int) (model.BookResponse, error)
}
