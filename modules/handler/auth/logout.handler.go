package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (authHandler *Handler) LogoutUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		cookie := &http.Cookie{
			Name:    "jwt",
			Value:   "",
			Expires: time.Now().Add(-time.Hour),
		}

		e.SetCookie(cookie)

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Logout successfull",
		})
	}
}
