package book

import (
	"errors"

	"github.com/eddiandrachen/Golang-DTS-challange-1/module/model"
)

type RepositoryBookImpl struct {
	Data map[uint64]model.Books
}

func NewBookRepository() RepositoryBook {
	DataStore := map[uint64]model.Books{
		0: {
			Id:     0,
			Author: "Tara",
			Title:  "Tau",
			//Price:       "sepuluh ribu",
			Description: "Buku",
			DeletedAt:   false,
		},
		1: {
			Id:     1,
			Author: "Bongs",
			Title:  "Beceq",
			//Price:       "RP20.000",
			Description: "serr geserr",
			DeletedAt:   false,
		},
		2: {
			Id:     2,
			Author: "Calman",
			Title:  "Peers",
			//Price:       "RP.12.000",
			Description: "Buku buah",
			DeletedAt:   false,
		},
	}
	return &RepositoryBookImpl{
		Data: DataStore,
	}
}

// b BookRepositoryImpl
func (b *RepositoryBookImpl) GetAllBook() ([]model.Books, error) {
	var data []model.Books
	for _, vals := range b.Data {
		data = append(data, vals)
	}
	return data, nil
}

func (b *RepositoryBookImpl) GetBookById(id int) (model.Books, error) {
	var data model.Books
	valid := false

	for _, vals := range b.Data {
		if id == int(vals.Id) && vals.DeletedAt != true {
			data = model.Books{
				Id:          vals.Id,
				Title:       vals.Title,
				Author:      vals.Author,
				Description: vals.Description,
			}
			valid = true
			break
		}
	}

	if valid {
		return data, nil
	} else {
		err := errors.New("Data tidak di temukan")
		return data, err
	}
}

func (b *RepositoryBookImpl) CreateBook(book model.Books) model.Books {
	counter := len(b.Data)
	book.Id = counter
	book.DeletedAt = false
	b.Data[uint64(counter)] = book
	return b.Data[uint64(counter)]
}

func (b *RepositoryBookImpl) UpdateBook(book model.Books) (model.Books, error) {
	if b.Data[uint64(book.Id)].DeletedAt {
		err := errors.New("Data tidak di temukan")
		return book, err
	} else {
		b.Data[uint64(book.Id)] = book
		return book, nil
	}
}

func (b *RepositoryBookImpl) DeleteBook(id int) error {

	if id < len(b.Data)-1 && b.Data[uint64(id)].DeletedAt == false {
		book := b.Data[uint64(id)]
		book.DeletedAt = true
		b.Data[uint64(id)] = book
		return nil
	} else {
		return errors.New("Data tidak di temukan")
	}
}
