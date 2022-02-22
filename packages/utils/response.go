package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"ginIMApi/constants/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type JsonStruct map[string]interface{}

// Response setting gin.JSON
func (g *Gin) Response(code int, data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
	return
}

func (g *Gin) ResponseWithHttp(httpCode int, code int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
	return
}

