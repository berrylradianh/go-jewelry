package products

import (
	"fmt"
	"net/http"
	"strconv"

	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (productCategoryHandler *Handler) GetAllProductCategories() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productCategories *[]ep.ProductCategory

		productCategories, err := productCategoryHandler.Usecase.GetAllProductCategories()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get All Product Categories",
			"users":   productCategories,
		})
	}
}

func (productCategoryHandler *Handler) GetProductCategoryById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productCategory *ep.ProductCategory
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		productCategory, err = productCategoryHandler.Usecase.GetProductCategoryById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get Product Category",
			"users":   productCategory,
		})
	}
}

func (productCategoryHandler *Handler) CreateProductCategory() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productCategory *ep.ProductCategory
		if err := e.Bind(&productCategory); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(productCategory); err != nil {
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

		err := productCategoryHandler.Usecase.CreateProductCategory(productCategory)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create Product Category",
		})
	}
}

func (productCategoryHandler *Handler) UpdateProductCategory() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productCategory *ep.ProductCategory
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		productCategory, err = productCategoryHandler.Usecase.GetProductCategoryById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&productCategory); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = productCategoryHandler.Usecase.UpdateProductCategory(int(productCategory.ID), productCategory)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Product Category",
		})
	}
}

func (productCategoryHandler *Handler) DeleteProductCategory() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productCategory *ep.ProductCategory
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		productCategory, err = productCategoryHandler.Usecase.GetProductCategoryById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = productCategoryHandler.Usecase.DeleteProductCategory(int(productCategory.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Product Category",
		})
	}
}
