package response

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	Data    interface{}
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Error       string
}
