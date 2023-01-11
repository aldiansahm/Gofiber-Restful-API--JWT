package response

import "time"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	Data    interface{}
}

type Book struct {
	Id          uint64
	Title       string
	Description string
	Author      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Rating      interface{}
}
