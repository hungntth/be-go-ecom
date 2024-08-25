package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code 	int 		`json:"code"` // status code
	Message string 		`json:"message"` // toats error
	Data 	interface{} `json:"data"` // data return
}

func SucessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: 		code,
		Message: 	msg[code],
		Data: 		data,
	})
}

func ErrorResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: 		code,
		Message: 	msg[code],
		Data: 		nil,
	})
}