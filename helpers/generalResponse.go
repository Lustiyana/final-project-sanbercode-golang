package helpers

import (
	"github.com/gin-gonic/gin"

	"simple-social-media/structs"
)

func GeneralResponse(ctx *gin.Context, statusCode int, success bool, message string, data interface{}, err interface{}) {
	var responseData interface{}

	if err != nil {
		responseData = err
	} else if data != nil {
		responseData = data
	} else {
		responseData = map[string]interface{}{}
	}
	
	response := structs.Response{
		Success:  success,
		Message: message,
		Data:    responseData,
	}

	ctx.JSON(statusCode, response)
}