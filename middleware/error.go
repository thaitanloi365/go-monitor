package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CustomErrorHandler error handler
func CustomErrorHandler(err error, c echo.Context) {
	var header = http.StatusInternalServerError
	var message interface{} = "Internal Server Error"
	var detail interface{} = err.Error()
	var code = header

	switch e := err.(type) {

	case *echo.HTTPError:
		header = e.Code
		code = header
		detail = fmt.Sprintf("[Echo Error] %s", e.Error())
		message = e.Message
		break

	}
	var reqID = c.Response().Header().Get(echo.HeaderXRequestID)

	var response = map[string]interface{}{
		"code":      code,
		"message":   message,
		"detail":    detail,
		"requestID": reqID,
	}

	c.JSON(header, response)
}
