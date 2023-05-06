package roles

import (
	"fmt"
	"net/http"
	"strconv"

	er "github.com/berrylradianh/go-jewelry/modules/entity/roles"
	ur "github.com/berrylradianh/go-jewelry/modules/usecase/roles"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase *ur.Usecase
}

func (roleHandler *Handler) GetAllRoles() echo.HandlerFunc {
	return func(e echo.Context) error {
		var roles *[]er.Role

		roles, err := roleHandler.Usecase.GetAllRoles()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get All Roles",
			"roles":   roles,
		})
	}
}

func (roleHandler *Handler) GetRoleById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var role *er.Role
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		role, err = roleHandler.Usecase.GetRoleById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get Role",
			"role":    role,
		})
	}
}

func (roleHandler *Handler) CreateRole() echo.HandlerFunc {
	return func(e echo.Context) error {
		var role *er.Role
		if err := e.Bind(&role); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(role); err != nil {
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

		err := roleHandler.Usecase.CreateRole(role)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create Role",
		})
	}
}

func (roleHandler *Handler) UpdateRole() echo.HandlerFunc {
	return func(e echo.Context) error {
		var role *er.Role
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		role, err = roleHandler.Usecase.GetRoleById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&role); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = roleHandler.Usecase.UpdateRole(int(role.ID), role)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Role",
		})
	}
}

func (roleHandler *Handler) DeleteRole() echo.HandlerFunc {
	return func(e echo.Context) error {
		var role *er.Role
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		role, err = roleHandler.Usecase.GetRoleById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = roleHandler.Usecase.DeleteRole(int(role.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Role",
		})
	}
}
