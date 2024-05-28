package exceptions

import (
	"encoding/json"
	"net/http"
)

type ErrorType struct {
	Code           string         `json:"code"`
	Type           string         `json:"type"`
	Forward        map[string]any `json:"forward,omitempty"`
	Abort          bool           `json:"abort"`
	Notify         bool           `json:"notify"`
	HttpStatusCode int            `json:"-"`
}

type Error struct {
	Messages []string `json:"messages"`
	ErrorType
}

func (e Error) Error() (message string) {
	l := len(e.Messages)
	for i, m := range e.Messages {
		message += m
		if l > 1 && l > i {
			message += " ; "
		}
	}
	return message
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
