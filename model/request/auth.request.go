package request

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ClaimsJWT struct {
	Expires int64
	Email   string
	Role    string
}
