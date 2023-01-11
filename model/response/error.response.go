package response

type ErrorResponse struct {
	FailedField string
	Tag         string
	Error       string
}
