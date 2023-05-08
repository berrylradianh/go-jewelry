package products

import (
	"fmt"
	"net/http"
	"strconv"

	ent "github.com/berrylradianh/go-jewelry/modules/entity"
	uc "github.com/berrylradianh/go-jewelry/modules/usecase/carts"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase *uc.Usecase
}

func (cartHandler *Handler) GetAllCarts() echo.HandlerFunc {
	return func(e echo.Context) error {
		var carts *[]ent.Cart

		carts, err := cartHandler.Usecase.GetAllCarts()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get All Carts",
			"carts":   carts,
		})
	}
}

func (cartHandler *Handler) GetCartById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var cart *ent.Cart
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		cart, err = cartHandler.Usecase.GetCartById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get Cart",
			"cart":    cart,
		})
	}
}

func (cartHandler *Handler) CreateCart() echo.HandlerFunc {
	return func(e echo.Context) error {
		var cart *ent.Cart
		if err := e.Bind(&cart); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(cart); err != nil {
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

		err := cartHandler.Usecase.CreateCart(cart)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create Cart",
		})
	}
}

func (cartHandler *Handler) UpdateCart() echo.HandlerFunc {
	return func(e echo.Context) error {
		var cart *ent.Cart
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		cart, err = cartHandler.Usecase.GetCartById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&cart); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = cartHandler.Usecase.UpdateCart(int(cart.ID), cart)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Cart",
		})
	}
}

func (cartHandler *Handler) DeleteCart() echo.HandlerFunc {
	return func(e echo.Context) error {
		var cart *ent.Cart
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		cart, err = cartHandler.Usecase.GetCartById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = cartHandler.Usecase.DeleteCart(int(cart.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Cart",
		})
	}
}
