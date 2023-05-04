package products

import (
	"fmt"
	"net/http"
	"strconv"

	ep "github.com/berrylradianh/go-jewelry/modules/entity/products"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (productMaterialHandler *Handler) GetAllProductMaterials() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productMaterials *[]ep.ProductMaterial

		productMaterials, err := productMaterialHandler.Usecase.GetAllProductMaterials()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":           "Success Get All Product Materials",
			"product_materials": productMaterials,
		})
	}
}

func (productMaterialHandler *Handler) GetProductMaterialById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productMaterial *ep.ProductMaterial
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		productMaterial, err = productMaterialHandler.Usecase.GetProductMaterialById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":          "Success Get Product Material",
			"product_material": productMaterial,
		})
	}
}

func (productMaterialHandler *Handler) CreateProductMaterial() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productMaterial *ep.ProductMaterial
		if err := e.Bind(&productMaterial); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(productMaterial); err != nil {
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

		err := productMaterialHandler.Usecase.CreateProductMaterial(productMaterial)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create Product Material",
		})
	}
}

func (productMaterialHandler *Handler) UpdateProductMaterial() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productMaterial *ep.ProductMaterial
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		productMaterial, err = productMaterialHandler.Usecase.GetProductMaterialById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&productMaterial); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = productMaterialHandler.Usecase.UpdateProductMaterial(int(productMaterial.ID), productMaterial)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Product Material",
		})
	}
}

func (productMaterialHandler *Handler) DeleteProductMaterial() echo.HandlerFunc {
	return func(e echo.Context) error {
		var productMaterial *ep.ProductMaterial
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		productMaterial, err = productMaterialHandler.Usecase.GetProductMaterialById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = productMaterialHandler.Usecase.DeleteProductMaterial(int(productMaterial.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Product Material",
		})
	}
}
