package middleware

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thaitanloi365/go-monitor/models"
)

// IsBasicAuth Authorization middleware
func IsBasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		var cc = c.(*models.CustomContext)

		var u = cc.Config.AdminUserName
		var p = cc.Config.AdminUserPassword

		if subtle.ConstantTimeCompare([]byte(username), []byte(u)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(p)) == 1 {
			return true, nil
		}

		return false, nil
	})

}
