package err

import (
	"fmt"
	"net/http"
)

const (
	httpConcatMessage        = "error code: %v. message: %v"
	messageDuplicatedIdError = "duplicate id"
	invalidRequest           = "invalid request"
	notFound                 = "Not Found"
)

var (
	DuplicatedIdError = &HttpError{
		Msg:  messageDuplicatedIdError,
		Code: http.StatusConflict,
	}

	BadRequestError = &HttpError{
		Msg:  invalidRequest,
		Code: http.StatusBadRequest,
	}

	NotFoundError = &HttpError{
		Msg:  notFound,
		Code: http.StatusNotFound,
	}

	IdParamCanNotBeEmptyError = &HttpError{
		Msg:  "id param can not be empty",
		Code: http.StatusBadRequest,
	}
)

type HttpError struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func (h HttpError) Error() string {
	return fmt.Sprintf(httpConcatMessage, h.Code, h.Msg)
}

func NewHttpError(msg string, code int) *HttpError {
	return &HttpError{
		Msg:  msg,
		Code: code,
	}
}
