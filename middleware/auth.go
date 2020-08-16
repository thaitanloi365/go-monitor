package middleware

import (
	"crypto/subtle"
	"fmt"
	"net/http"

	"github.com/thaitanloi365/go-monitor/config"
	"github.com/thaitanloi365/go-monitor/errs"
	"github.com/thaitanloi365/go-monitor/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const contextKey = "user_token"

// IsAuthorized Authorization middleware
func IsAuthorized() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:       &models.JwtClaims{},
		ContextKey:   contextKey,
		SigningKey:   []byte(config.GetInstance().JWTSecret),
		ErrorHandler: transformError,
	})
}

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

// CheckTokenExpiredAndAttachUserInfo Check token expired middleware
func CheckTokenExpiredAndAttachUserInfo() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var cc = c.(*models.CustomContext)

			jwtClaims, err := cc.GetJwtClaims()
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			}

			var user models.User
			err = cc.DB.First(&user, "id = ?", jwtClaims.ID).Error
			if err != nil {
				return errs.ErrTokenInvalid
			}
			var isValid = jwtClaims.VerifyIssuer(user.TokenIssuer, true)
			if !isValid {
				isValid = jwtClaims.VerifyIssuer("admin", true)
			}

			if !isValid {
				return errs.ErrTokenInvalid
			}

			c.Set("user", user)

			return next(c)
		}
	}

}

// RequireRoles Authorization middleware
func RequireRoles(next echo.HandlerFunc, roles ...string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cc = c.(*models.CustomContext)
		var isValidRole bool = false

		jwtClaims, err := cc.GetJwtClaims()
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		for _, role := range roles {
			if role == jwtClaims.Audience {
				isValidRole = true
				break
			}
		}

		if isValidRole == false {
			return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("This route only valid for rules: %v", roles))
		}

		return next(c)
	}
}

func transformError(err error) error {
	if err == middleware.ErrJWTMissing {
		return errs.ErrTokenMissing
	}

	if e, ok := err.(*echo.HTTPError); ok {
		if e.Code == http.StatusBadRequest {
			return errs.ErrTokenMissing
		}

		if e.Code == http.StatusUnauthorized {
			return errs.ErrTokenInvalid
		}
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errs.ErrTokenInvalid
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return errs.ErrTokenExpired
		} else {
			return errs.ErrTokenInvalid
		}
	}

	return err
}
