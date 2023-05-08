package users

import (
	"fmt"
	"net/http"
	"strconv"

	ent "github.com/berrylradianh/go-jewelry/modules/entity"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (UserDetailHandler *Handler) GetAllUserDetails() echo.HandlerFunc {
	return func(e echo.Context) error {
		var userDetails *[]ent.UserDetail

		userDetails, err := UserDetailHandler.Usecase.GetAllUserDetails()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":      "Success Get All User Details",
			"user_details": userDetails,
		})
	}
}

func (UserDetailHandler *Handler) GetUserDetailById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var userDetail *ent.UserDetail
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		userDetail, err = UserDetailHandler.Usecase.GetUserDetailById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":     "Success Get User Detail",
			"user_detail": userDetail,
		})
	}
}

func (UserDetailHandler *Handler) CreateUserDetail() echo.HandlerFunc {
	return func(e echo.Context) error {
		var userDetail *ent.UserDetail
		if err := e.Bind(&userDetail); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(userDetail); err != nil {
			message := ""
			for _, e := range err.(validator.ValidationErrors) {
				if e.Tag() == "required" {
					message = fmt.Sprintf("%s is required ", e.Field())
				}
			}
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
				// "errors":  err.Error(),
			})
		}

		err := UserDetailHandler.Usecase.CreateUserDetail(userDetail)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create User Detail",
		})
	}
}

func (UserDetailHandler *Handler) UpdateUserDetail() echo.HandlerFunc {
	return func(e echo.Context) error {
		var userDetail *ent.UserDetail
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		userDetail, err = UserDetailHandler.Usecase.GetUserDetailById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&userDetail); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = UserDetailHandler.Usecase.UpdateUserDetail(int(userDetail.ID), userDetail)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update User Detail",
		})
	}
}

func (UserDetailHandler *Handler) DeleteUserDetail() echo.HandlerFunc {
	return func(e echo.Context) error {
		var userDetail *ent.UserDetail
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		userDetail, err = UserDetailHandler.Usecase.GetUserDetailById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = UserDetailHandler.Usecase.DeleteUserDetail(int(userDetail.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete User Detail",
		})
	}
}
