package model

type Book struct {
	Id     int    `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	//Price       string `json:"Price"`
	Description string `json:"Description"`
	//CreatedAt   time.Time      `json:"created_at"`
	//Update      time.Time      `json:"updated_at"`
	DeletedAt bool
}
