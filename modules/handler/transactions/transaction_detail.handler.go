package transactions

import (
	"fmt"
	"net/http"
	"strconv"

	ent "github.com/berrylradianh/go-jewelry/modules/entity"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (transactionDetailHandler *Handler) GetAllTransactionDetails() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transactionDetails *[]ent.TransactionDetail

		transactionDetails, err := transactionDetailHandler.Usecase.GetAllTransactionDetails()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":             "Success Get All Transaction Details",
			"transaction_details": transactionDetails,
		})
	}
}

func (transactionDetailHandler *Handler) GetTransactionDetailById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transactionDetail *ent.TransactionDetail
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		transactionDetail, err = transactionDetailHandler.Usecase.GetTransactionDetailById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":            "Success Get Transaction Detail",
			"transaction_detail": transactionDetail,
		})
	}
}

func (transactionDetailHandler *Handler) CreateTransactionDetail() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transactionDetail *ent.TransactionDetail
		if err := e.Bind(&transactionDetail); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(transactionDetail); err != nil {
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

		product, err := transactionDetailHandler.Usecase.FindProduct(int(transactionDetail.Product_id))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		transactionDetail.Price = product.Price * float64(transactionDetail.Qty)

		err = transactionDetailHandler.Usecase.CreateTransactionDetail(transactionDetail)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create Transaction Detail",
		})
	}
}

func (transactionDetailHandler *Handler) UpdateTransactionDetail() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transactionDetail *ent.TransactionDetail
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		transactionDetail, err = transactionDetailHandler.Usecase.GetTransactionDetailById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&transactionDetail); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = transactionDetailHandler.Usecase.UpdateTransactionDetail(int(transactionDetail.ID), transactionDetail)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Transaction Detail",
		})
	}
}

func (transactionDetailHandler *Handler) DeleteTransactionDetail() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transactionDetail *ent.TransactionDetail
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		transactionDetail, err = transactionDetailHandler.Usecase.GetTransactionDetailById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = transactionDetailHandler.Usecase.DeleteTransactionDetail(int(transactionDetail.ID))
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
