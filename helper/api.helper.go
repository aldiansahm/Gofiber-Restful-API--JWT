package helper

import (
	"errors"
	"github.com/aldiansahm7654/go-restapi-fiber/model/response"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetResponse(status int, data interface{}, err error) response.Response {
	var response response.Response

	switch status {
	case 200:
		response.Message = "Success"
	default:
		response.Message = err.Error()
	}

	response.Status = status
	response.Data = data

	return response
}

func GetErrorValidate(data interface{}) []*response.ErrorResponse {
	var errors []*response.ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Error = err.Error()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateStruct(data interface{}) (response.Response, error) {
	var res response.Response
	errValidate := GetErrorValidate(data)
	if errValidate != nil {
		err := errors.New("validation errors")
		res = GetResponse(400, errValidate, err)
		return res, err
	}
	return res, nil
}
