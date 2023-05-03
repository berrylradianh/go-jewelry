package products

import (
	"fmt"
	"net/http"
	"strconv"

	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"
	up "github.com/berrylradianh/go-jewelry/modules/usecase/products"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase *up.Usecase
}

func (productHandler *Handler) GetAllProducts() echo.HandlerFunc {
	return func(e echo.Context) error {
		var products *[]ep.Product

		products, err := productHandler.Usecase.GetAllProducts()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get All Products",
			"users":   products,
		})
	}
}

func (productHandler *Handler) GetProductById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var product *ep.Product
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		product, err = productHandler.Usecase.GetProductById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get Product",
			"users":   product,
		})
	}
}

func (productHandler *Handler) CreateProduct() echo.HandlerFunc {
	return func(e echo.Context) error {
		var product *ep.Product
		if err := e.Bind(&product); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(product); err != nil {
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

		err := productHandler.Usecase.CreateProduct(product)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create Product",
		})
	}
}

func (productHandler *Handler) UpdateProduct() echo.HandlerFunc {
	return func(e echo.Context) error {
		var product *ep.Product
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		product, err = productHandler.Usecase.GetProductById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&product); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		result := productHandler.Usecase.UpdateProduct(int(product.ID), product)
		if result == 0 {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Product",
		})
	}
}

func (productHandler *Handler) DeleteProduct() echo.HandlerFunc {
	return func(e echo.Context) error {
		var product *ep.Product
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		product, err = productHandler.Usecase.GetProductById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = productHandler.Usecase.DeleteProduct(int(product.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Product",
		})
	}
}
