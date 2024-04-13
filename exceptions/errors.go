package exceptions

import (
	"encoding/json"
	"net/http"
)

type ErrorType struct {
	Code           string `json:"code"`
	Type           string `json:"type"`
	HttpStatusCode int    `json:"-"`
}

type Error struct {
	Message string `json:"message"`
	ErrorType
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func (e Error) WriteHttp(writer http.ResponseWriter) (err error) {
	b, err := e.ToJSON()
	writer.WriteHeader(e.HttpStatusCode)
	_, err = writer.Write(b)
	return
}
