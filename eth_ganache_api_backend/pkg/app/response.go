package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api_backend/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
// func (g *Gin) Response(httpCode, errCode int, data interface{}) {
// g.C.JSON(httpCode, Response{
// Code: errCode,
// Msg:  e.GetMsg(errCode),
// Data: data,
// })
// return
// }

func Json(c *gin.Context, httpCode, code int, data interface{}) {
	c.JSON(httpCode, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}

func JsonSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}

func JsonError(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}

func JsonBadRequest(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}
