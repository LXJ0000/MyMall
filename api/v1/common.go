package v1

import (
	"MyMall/serializer"
	"github.com/goccy/go-json"
	"net/http"
)

func ErrorResponse(err error) serializer.Response {
	msg := "参数错误"
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		msg = "JSON类型不匹配"
	}
	return serializer.Response{
		Status: http.StatusBadRequest,
		Msg:    msg,
		Error:  err.Error(),
	}
}
