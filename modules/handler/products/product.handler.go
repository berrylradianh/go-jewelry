package products

import (
	"fmt"
	"net/http"
	"strconv"

	ent "github.com/berrylradianh/go-jewelry/modules/entity"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (productHandler *Handler) GetAllProducts() echo.HandlerFunc {
	return func(e echo.Context) error {
		var products *[]ent.Product

		products, err := productHandler.Usecase.GetAllProducts()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Success Get All Products",
			"products": products,
		})
	}
}

func (productHandler *Handler) GetProductById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var product *ent.Product
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
			"product": product,
		})
	}
}

func (productHandler *Handler) CreateProduct() echo.HandlerFunc {
	return func(e echo.Context) error {
		var product *ent.Product
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
		var product *ent.Product
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

		err = productHandler.Usecase.UpdateProduct(int(product.ID), product)
		if err != nil {
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
		var product *ent.Product
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

func (productHandler *Handler) SortProducts() echo.HandlerFunc {
	return func(e echo.Context) error {
		var products *[]ent.Product
		sortBy := e.QueryParam("sort")
		sortOrder := e.QueryParam("order")

		validSortFields := []string{"name", "created_at"}
		isValidSortField := false
		for _, field := range validSortFields {
			if sortBy == field {
				isValidSortField = true
				break
			}
		}

		if !isValidSortField {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid sort parameter",
			})
		}

		if sortOrder != "asc" && sortOrder != "desc" {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid order parameter",
			})
		}

		products, err := productHandler.Usecase.SortProducts(sortBy, sortOrder)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to sort products",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Success Get All Products",
			"products": products,
		})
	}
}

func (productHandler *Handler) FilterProducts() echo.HandlerFunc {
	return func(e echo.Context) error {
		var products *[]ent.Product
		var err error

		param := e.QueryParam("param")
		switch param {
		case "material":
			material := e.QueryParam("material")
			products, err = productHandler.Usecase.FilterProductsByMaterial(material)
		case "category":
			category := e.QueryParam("category")
			products, err = productHandler.Usecase.FilterProductsByCategory(category)
		default:
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Invalid filter parameter",
			})
		}

		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Success Get All Products",
			"products": products,
		})
	}
}

func (productHandler *Handler) SearchProducts() echo.HandlerFunc {
	return func(e echo.Context) error {
		var products *[]ent.Product
		var err error

		param := e.QueryParam("param")
		switch param {
		case "name":
			productName := e.QueryParam("name")
			if productName == "" {
				return e.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "Product name is required",
				})
			}
			products, err = productHandler.Usecase.SearchProductsByName(productName)
		case "category":
			productCategory := e.QueryParam("category")
			if productCategory == "" {
				return e.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "Product category is required",
				})
			}
			products, err = productHandler.Usecase.SearchProductsByCategory(productCategory)
		default:
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Invalid filter parameter",
			})
		}

		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Success Get All Products",
			"products": products,
		})
	}
}
