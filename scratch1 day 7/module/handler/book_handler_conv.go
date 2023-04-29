package handler

import "github.com/eddiandrachen/Golang-DTS-challange-1/module/model"

func BookDomainToResp(data model.Book) model.BookResponse {
	resp := model.BookResponse{
		Id:          data.Id,
		Title:       data.Title,
		Author:      data.Author,
		Description: data.Description,
	}
	return resp
}
