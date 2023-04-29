package book

import (
	"github.com/eddiandrachen/Golang-DTS-challange-1/module/model"
)

type RepositoryBook interface {
	GetAllBook([]model.Book, error)
	GetByID(Id int) (model.Book, error)
	CreateBook(model.Book) model.Book
	Update(model.Book) (model.Book, error)
	DeletedAt(id int) error
}
