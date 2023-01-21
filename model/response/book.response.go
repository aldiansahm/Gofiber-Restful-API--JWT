package response

import "time"

type Book struct {
	Id          uint64      `json:"id"`
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	Author      *string     `json:"author"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Rating      interface{} `json:"rating"`
}

type BookRating struct {
	Id          uint64  `json:"id,omitempty" gorm:"primaryKey,column:id"`
	BookId      *string `json:"id_buku, omitempty" gorm:"index"`
	Picture     *string `json:"picture"`
	Description *string `json:"description"`
	Rating      int     `json:"rating"`
}
