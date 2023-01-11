package helper

import (
	"errors"
	"github.com/aldiansahm7654/go-restapi-fiber/model/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateTokenJWT(user request.ClaimsJWT) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["role"] = user.Role
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix() // a week

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func indexOf(element string, data []string) (int, error) {
	for k, v := range data {
		if element == v {
			return k, nil
			break
		}
	}
	err := errors.New("unauthorized access")
	return -1, err //not found.
}

func ExtractTokenMetadata(c *fiber.Ctx, Role ...string) (*request.ClaimsJWT, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		expires := int64(claims["exp"].(float64))

		if len(Role) > 0 {
			if _, errRole := indexOf(claims["role"].(string), Role); errRole != nil {
				return nil, errRole
			}
		}

		return &request.ClaimsJWT{
			Expires: expires,
			Email:   claims["email"].(string),
			Role:    claims["role"].(string),
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
