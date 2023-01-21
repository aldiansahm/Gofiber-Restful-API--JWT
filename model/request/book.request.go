package request

type Book struct {
	Title       string `json:"title" validate:"required,min=3,max=32" `
	Description string `json:"description" validate:"required"`
	Author      string `json:"author" validate:"required"`
}
