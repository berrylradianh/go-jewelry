package auth

import (
	"fmt"
	"net/http"

	ent "github.com/berrylradianh/go-jewelry/modules/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (authHandler *Handler) RegisterUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user *ent.User
		if err := e.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(user); err != nil {
			message := ""
			for _, e := range err.(validator.ValidationErrors) {
				if e.Tag() == "required" {
					message = fmt.Sprintf("%s is required ", e.Field())
				} else if e.Tag() == "email" {
					message = "Invalid email address "
				}
			}
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
				// "errors":  err.Error(),
			})
		}

		err := authHandler.Usecase.RegisterUser(user)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Register Successfull",
		})
	}
}
