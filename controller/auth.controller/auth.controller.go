package auth_controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aldiansahm7654/go-restapi-fiber/config/database"
	"github.com/aldiansahm7654/go-restapi-fiber/helper"
	"github.com/aldiansahm7654/go-restapi-fiber/model/entity"
	"github.com/aldiansahm7654/go-restapi-fiber/model/request"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req request.Login
	var data entity.User

	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	db := database.DBConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT name, email, address, phone, role, created_at, updated_at " +
		"FROM user WHERE email=? AND password=?"
	result := db.QueryRowContext(ctx, query, req.Email, req.Password).Scan(&data.Name,
		&data.Email, &data.Address, &data.Phone, &data.Role, &data.CreatedAt, &data.UpdatedAt)
	switch {
	case result == sql.ErrNoRows:
		errResult := errors.New("unauthorized")
		res := helper.GetResponse(fiber.StatusUnauthorized, nil, errResult)
		return c.Status(res.Status).JSON(res)
	case result != nil:
		errResult := errors.New("bad credentials")
		res := helper.GetResponse(500, nil, errResult)
		return c.Status(res.Status).JSON(res)
	}

	claimsJWT := request.ClaimsJWT{
		Email: data.Email,
		Role:  data.Role,
	}
	token, err := helper.GenerateTokenJWT(claimsJWT)
	if err != nil {
		res := helper.GetResponse(400, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	res := helper.GetResponse(200, data, err)
	res.Token = token
	return c.Status(res.Status).JSON(res)
}
