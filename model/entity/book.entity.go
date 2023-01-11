package entity

import "time"

type Book struct {
	Id          uint64      `json:"id,omitempty" gorm:"primaryKey,column:id"`
	Title       string      `json:"title" gorm:"type:varchar(255)"`
	Description string      `json:"description" gorm:"type:text"`
	Author      string      `json:"author" gorm:"type:varchar(255)"`
	CreatedAt   time.Time   `json:"created_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	BookRating  *BookRating `json:",omitempty"`
}

type BookRating struct {
	Id          uint64 `json:"id,omitempty" gorm:"primaryKey,column:id"`
	BookId      string `json:"id_buku, omitempty" gorm:"index"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}

func (Book) TableName() string {
	return "book"
}

func (BookRating) TableName() string {
	return "book_rating"
}
