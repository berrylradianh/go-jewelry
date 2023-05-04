package products

import (
	"fmt"
	"net/http"
	"strconv"

	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (productDescriptionHandler *Handler) GetAllProductDescriptions() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productDescriptions *[]ep.ProductDescription

		productDescriptions, err := productDescriptionHandler.Usecase.GetAllProductDescriptions()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":              "Success Get All Product Descriptions",
			"product_descriptions": productDescriptions,
		})
	}
}

func (productDescriptionHandler *Handler) GetProductDescriptionById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productDescription *ep.ProductDescription
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		productDescription, err = productDescriptionHandler.Usecase.GetProductDescriptionById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":             "Success Get Product Description",
			"product_description": productDescription,
		})
	}
}

func (productDescriptionHandler *Handler) CreateProductDescription() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productDescription *ep.ProductDescription
		if err := e.Bind(&productDescription); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(productDescription); err != nil {
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

		err := productDescriptionHandler.Usecase.CreateProductDescription(productDescription)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create Product Description",
		})
	}
}

func (productDescriptionHandler *Handler) UpdateProductDescription() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productDescription *ep.ProductDescription
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		productDescription, err = productDescriptionHandler.Usecase.GetProductDescriptionById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&productDescription); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = productDescriptionHandler.Usecase.UpdateProductDescription(int(productDescription.ID), productDescription)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Product Description",
		})
	}
}
