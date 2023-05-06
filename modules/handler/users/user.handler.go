package users

import (
	"fmt"
	"net/http"
	"strconv"

	eu "github.com/berrylradianh/go-jewelry/modules/entity/users"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (UserHandler *Handler) GetAllUsers() echo.HandlerFunc {
	return func(e echo.Context) error {
		var users *[]eu.User

		users, err := UserHandler.Usecase.GetAllUsers()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get All Users",
			"users":   users,
		})
	}
}

func (UserHandler *Handler) GetUserById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user *eu.User
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		user, err = UserHandler.Usecase.GetUserById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get User",
			"user":    user,
		})
	}
}

func (UserHandler *Handler) CreateUser() echo.HandlerFunc {
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

		err := UserHandler.Usecase.CreateUser(user)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create User",
		})
	}
}

func (UserHandler *Handler) UpdateUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user *eu.User
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		user, err = UserHandler.Usecase.GetUserById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = UserHandler.Usecase.UpdateUser(int(user.ID), user)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update User",
		})
	}
}

func (UserHandler *Handler) DeleteUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user *eu.User
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		user, err = UserHandler.Usecase.GetUserById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = UserHandler.Usecase.DeleteUser(int(user.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete User",
		})
	}
}
