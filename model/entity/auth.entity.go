package entity

import (
	"time"
)

type User struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(151)"`
	Email     string    `json:"email" gorm:"type:varchar(100)"`
	Password  string    `json:"-" gorm:"column:password" gorm:"type:text"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone" gorm:"type:varchar(20)"`
	Role      string    `json:"role" gorm:"type:varchar(15)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "user"
}
