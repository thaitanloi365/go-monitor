package middleware

import (
	"fmt"
	"net/http"

	"github.com/thaitanloi365/go-monitor/errs"
	"github.com/thaitanloi365/go-monitor/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// CustomErrorHandler error handler
func CustomErrorHandler(err error, c echo.Context) {
	var header = http.StatusInternalServerError
	var message interface{} = "Internal Server Error"
	var detail interface{} = err.Error()
	var code = header

	switch e := err.(type) {
	case *validator.InvalidValidationError:
		header = http.StatusUnprocessableEntity
		code = header
		detail = fmt.Sprintf("[Invalid Validation Error] %s", e.Error())
		message = e.Error()

		break
	case validator.ValidationErrors:
		header = http.StatusUnprocessableEntity
		code = header
		detail = fmt.Sprintf("[Validation Error] %s", e.Error())
		message = "Unprocessable Entity"
		break
	case *echo.HTTPError:
		header = e.Code
		code = header
		detail = fmt.Sprintf("[Echo Error] %s", e.Error())
		message = e.Message
		break
	case *pq.Error:
		header = http.StatusInternalServerError
		code = header
		detail = fmt.Sprintf("[Database Error - %s] %s. %s", e.Code, e.Error(), e.Detail)
		message = "Internal Server Error"

		break
	case *errs.Error:
		header = e.Header
		code = e.Code
		detail = fmt.Sprintf("[Readable Error - %d] %s", e.Code, e.Error())
		message = e.Error()
		break
	}
	var reqID = c.Response().Header().Get(echo.HeaderXRequestID)

	var response = models.M{
		"code":      code,
		"message":   message,
		"detail":    detail,
		"requestID": reqID,
	}

	c.JSON(header, response)
}
