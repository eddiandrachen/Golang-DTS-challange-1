package handler

import "github.com/eddiandrachen/Golang-DTS-challange-1/module/model"

func BookResponse(data model.Book) model.Book {
	response := model.Book{
		Id:     data.Id,
		Title:  data.Title,
		Author: data.Author,
		//Price:       data.Price,
		Description: data.Description,
	}
	return response
}
