package BookController

import (
	"context"
	"errors"
	"fmt"
	"github.com/aldiansahm7654/go-restapi-fiber/config/database"
	"github.com/aldiansahm7654/go-restapi-fiber/helper"
	"github.com/aldiansahm7654/go-restapi-fiber/model/entity"
	"github.com/aldiansahm7654/go-restapi-fiber/model/request"
	"github.com/aldiansahm7654/go-restapi-fiber/model/response"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Index(c *fiber.Ctx) error {

	claims, err := helper.ExtractTokenMetadata(c, "admin")
	if err != nil {
		// Return status 500 and JWT parse error.
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	fmt.Println(claims)

	var books []response.Book
	var book response.Book

	db := database.DBConnection()
	defer db.Close()

	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	//defer cancel()

	query := "SELECT id, title, description, author, created_at, updated_at " +
		"FROM book"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()
	//fmt.Println(rows)
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Description, &book.Author, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}

		queries, _ := db.QueryContext(ctx, "SELECT id, book_id, picture, description, rating "+
			"FROM book_rating WHERE book_id = ?", book.Id)

		var ratings []entity.BookRating
		var rating entity.BookRating
		for queries.Next() {
			err := queries.Scan(&rating.Id, &rating.BookId, &rating.Picture, &rating.Description, &rating.Rating)
			if err != nil {
				res := helper.GetResponse(500, nil, err)
				return c.Status(res.Status).JSON(res)
			}
			ratings = append(ratings, rating)
		}
		queries.Close()
		book.Rating = ratings
		books = append(books, book)
	}

	res := helper.GetResponse(200, books, nil)
	return c.JSON(res)
}

func Create(c *fiber.Ctx) error {
	db := database.DBConnection()
	defer db.Close()

	var req request.Book

	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	resValidate, errValidate := helper.ValidateStruct(req)
	if errValidate != nil {
		return c.Status(resValidate.Status).JSON(resValidate)
	}

	data := entity.Book{
		Title:       req.Title,
		Description: req.Author,
		Author:      req.Author,
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_, err := db.ExecContext(ctx, "INSERT INTO book(title, description, author) VALUES(?, ?, ?)", data.Title, data.Description, data.Author)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	res := helper.GetResponse(200, req, nil)
	return c.JSON(res)
}

func Update(c *fiber.Ctx) error {
	db := database.DBConnection()
	defer db.Close()

	var req request.Book

	getId := c.Params("id")

	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(400, req, err)
		return c.Status(res.Status).JSON(res)
	}

	data := entity.Book{
		Title:       req.Title,
		Description: req.Author,
		Author:      req.Author,
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	row, err := db.ExecContext(ctx, "UPDATE book SET title =?,  description= ?,  author= ? WHERE id=?", data.Title, data.Description, data.Author, getId)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	result, err := row.RowsAffected()
	if err != nil {
		res := helper.GetResponse(400, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	if result < 1 {
		err := errors.New("Tidak ada data yang diupdate.")
		res := helper.GetResponse(400, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	idBook, _ := strconv.Atoi(getId)
	data.Id = uint64(idBook)

	res := helper.GetResponse(200, data, err)
	return c.JSON(res)
}

func Delete(c *fiber.Ctx) error {
	db := database.DBConnection()
	defer db.Close()
	getId := c.Params("id")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	row, err := db.ExecContext(ctx, "DELETE from book WHERE id=?", getId)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	result, err := row.RowsAffected()
	if err != nil {
		res := helper.GetResponse(400, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	if result < 1 {
		err := errors.New("Data tidak ditemukan.")
		res := helper.GetResponse(404, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	res := helper.GetResponse(200, nil, nil)
	return c.JSON(res)
}
