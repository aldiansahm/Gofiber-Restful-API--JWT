package response

type User struct {
	ID      uint    `json:"id,omitempty"`
	Name    *string `json:"name"`
	Email   string  `json:"email"`
	Address *string `json:"address"`
	Phone   *string `json:"phone"`
	Role    string  `json:"role"`
}
