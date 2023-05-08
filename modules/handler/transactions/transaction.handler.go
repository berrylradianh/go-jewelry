package transactions

import (
	"fmt"
	"net/http"
	"strconv"

	ent "github.com/berrylradianh/go-jewelry/modules/entity"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (transactionHandler *Handler) GetAllTransactions() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transactions *[]ent.Transaction

		transactions, err := transactionHandler.Usecase.GetAllTransactions()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":      "Success Get All Transactions",
			"transactions": transactions,
		})
	}
}

func (transactionHandler *Handler) GetTransactionById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transaction *ent.Transaction
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		transaction, err = transactionHandler.Usecase.GetTransactionById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":     "Success Get Transaction",
			"transaction": transaction,
		})
	}
}

func (transactionHandler *Handler) CreateTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transaction *ent.Transaction
		if err := e.Bind(&transaction); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(transaction); err != nil {
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

		err := transactionHandler.Usecase.CreateTransaction(transaction)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create Transaction",
		})
	}
}

func (transactionHandler *Handler) UpdateTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transaction *ent.Transaction
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		transaction, err = transactionHandler.Usecase.GetTransactionById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&transaction); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = transactionHandler.Usecase.UpdateTransaction(int(transaction.ID), transaction)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Transaction",
		})
	}
}

func (transactionHandler *Handler) DeleteTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		var transaction *ent.Transaction
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		transaction, err = transactionHandler.Usecase.GetTransactionById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = transactionHandler.Usecase.DeleteTransaction(int(transaction.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Transaction",
		})
	}
}
