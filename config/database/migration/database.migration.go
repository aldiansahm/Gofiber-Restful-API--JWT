package migration

import (
	"github.com/aldiansahm7654/go-restapi-fiber/model/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func init() {
	log.Println("Start Migration Database ...")
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/gofiber-restapi"))
	if err != nil {
		panic(err)
	}
	errBook := db.AutoMigrate(&entity.Book{}, &entity.User{}, &entity.BookRating{})
	if errBook != nil {
		log.Println(errBook)
	}
	log.Println("Database Migrated")

	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()
}
