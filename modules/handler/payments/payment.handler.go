package payments

import (
	"fmt"
	"net/http"
	"strconv"

	ent "github.com/berrylradianh/go-jewelry/modules/entity"
	up "github.com/berrylradianh/go-jewelry/modules/usecase/payments"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase *up.Usecase
}

func (paymentHandler *Handler) GetAllPayments() echo.HandlerFunc {
	return func(e echo.Context) error {
		var payments *[]ent.Payment

		payments, err := paymentHandler.Usecase.GetAllPayments()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Success Get All Payments",
			"payments": payments,
		})
	}
}

func (paymentHandler *Handler) GetPaymentById() echo.HandlerFunc {
	return func(e echo.Context) error {
		var payment *ent.Payment
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		payment, err = paymentHandler.Usecase.GetPaymentById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get Payment",
			"payment": payment,
		})
	}
}

func (paymentHandler *Handler) CreatePayment() echo.HandlerFunc {
	return func(e echo.Context) error {
		var payment *ent.Payment
		if err := e.Bind(&payment); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		if err := e.Validate(payment); err != nil {
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

		err := paymentHandler.Usecase.CreatePayment(payment)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Create Payments",
		})
	}
}

func (paymentHandler *Handler) UpdatePayment() echo.HandlerFunc {
	return func(e echo.Context) error {
		var payment *ent.Payment
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		payment, err = paymentHandler.Usecase.GetPaymentById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&payment); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid Request Body",
				// "errors":  err.Error(),
			})
		}

		err = paymentHandler.Usecase.UpdatePayment(int(payment.ID), payment)
		if err != nil {
			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "Nothing Updated",
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Payment",
		})
	}
}

func (paymentHandler *Handler) DeletePayment() echo.HandlerFunc {
	return func(e echo.Context) error {
		var payment *ent.Payment
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		payment, err = paymentHandler.Usecase.GetPaymentById(id)
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = paymentHandler.Usecase.DeletePayment(int(payment.ID))
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Payment",
		})
	}
}
