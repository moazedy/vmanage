package errorx

import (
	"errors"
	"net/http"
)

type ErrorX struct {
	EmbedError     error
	HttpStatusCode int
}

func NewErrorX(msg string, code int) ErrorX {
	return ErrorX{
		EmbedError:     errors.New(msg),
		HttpStatusCode: code,
	}
}

func NewErrorXWithError(err error, code int) ErrorX {
	return ErrorX{
		EmbedError:     err,
		HttpStatusCode: code,
	}
}

func NewInternalErrorX(err error) ErrorX {
	return ErrorX{
		EmbedError:     err,
		HttpStatusCode: http.StatusInternalServerError,
	}
}

func NewBadRequestErrorX(err error) ErrorX {
	return ErrorX{
		EmbedError:     err,
		HttpStatusCode: http.StatusBadRequest,
	}
}

func NewNotFoundError(err error) ErrorX {
	return ErrorX{
		EmbedError:     err,
		HttpStatusCode: http.StatusNotFound,
	}
}

func (ex ErrorX) IsNil() bool {
	return ex.EmbedError == nil
}

func (ex ErrorX) Error() string {
	return ex.EmbedError.Error()
}

func (ex ErrorX) Nil() ErrorX {
	return ErrorX{}
}

func (ex ErrorX) GetHttpStatusCode() int {
	return ex.HttpStatusCode
}

func (ex ErrorX) GetError() error {
	return ex.EmbedError
}
