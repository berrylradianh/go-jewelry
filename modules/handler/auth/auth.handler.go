package auth

import (
	"fmt"
	"net/http"
	"time"

	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	ra "github.com/berrylradianh/go-jewelry/modules/response/auth"
	ua "github.com/berrylradianh/go-jewelry/modules/usecase/auth"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase *ua.Usecase
}

func (authHandler *Handler) RegisterUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user *eu.User
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

		hashedPassword, err := authHandler.Usecase.RegisterUser(user)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		user.Password = string(hashedPassword)

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Register Successfull",
		})
	}
}

func (authHandler *Handler) LoginUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user *eu.User
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

		user, token, err := authHandler.Usecase.LoginUser(user.Email, user.Password)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
		}

		e.SetCookie(cookie)

		authResponse := ra.AuthResponse{
			ID:    int(user.ID),
			Email: user.Email,
			Token: token,
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login successfull",
			"users":   authResponse,
		})
	}
}

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
