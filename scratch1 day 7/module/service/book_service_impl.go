package service

import (
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/handler"
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/model"
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/repository/book"

	"github.com/gin-gonic/gin"
)

type BookServiceImpl struct {
	RepositoryBook book.RepositoryBook
}

func NewBookService(BookRepo book.RepositoryBook) BookService {
	return &BookServiceImpl{
		RepositoryBook: BookRepo,
	}
	//return
}

func (bs *BookServiceImpl) Insert(ctx *gin.Context, request model.BookCreateRequest) model.BookResponse {
	var data model.Book
	data.Author = request.Author
	data.Description = request.Description
	data.Title = request.Title
	data = bs.RepositoryBook.CreateBook(data)
	return handler.BookDomainToResp(data)
}

func (bs *BookServiceImpl) Update(ctx *gin.Context, request model.BookUpdateRequest) (model.BookResponse, error) {
	var data model.Book
	data.Author = request.Author
	data.Description = request.Description
	data.Title = request.Title
	data.Id = request.Id

	_, err := bs.RepositoryBook.GetBookById(data.Id)
	if err != nil {
		return handler.BookDomainToResp(data), err
	}

	data, err = bs.RepositoryBook.UpdateBook(data)
	if err != nil {
		return handler.BookDomainToResp(data), err
	}

	return handler.BookDomainToResp(data), nil
}

func (bs *BookServiceImpl) Delete(ctx *gin.Context, id int) error {
	err := bs.RepositoryBook.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookServiceImpl) GetAll(ctx *gin.Context) []model.BookResponse {
	data, err := bs.RepositoryBook.GetAllBook()
	if err != nil {
		panic(err)
	}

	var resp []model.BookResponse
	for _, vals := range data {
		resp = append(resp, handler.BookDomainToResp(vals))
	}
	return resp
}

func (bs *BookServiceImpl) GetById(ctx *gin.Context, id int) (model.BookResponse, error) {

	data, err := bs.RepositoryBook.GetBookById(id)

	if err != nil {
		return handler.BookDomainToResp(data), err
	}

	return handler.BookDomainToResp(data), err
}
