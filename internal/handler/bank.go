package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
	
	"moneytransfer/domain"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type BankService interface {
	GetAccountByNumberAndName(ctx context.Context, accountNumber string, accountName string) (domain.Account, error)
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) error
	DisburseTransaction(ctx context.Context, refID string) error
}

type BankHandler struct {
	bankService BankService
}

func NewBankHandler(e *echo.Echo, bankService BankService) {
	handler := &BankHandler{
		bankService: bankService,
	}
	e.GET("/api/v1/bank/account/:accountno/:accountname", handler.GetAccount)
	e.POST("/api/v1/bank/transaction", handler.CreateTransaction)
	e.PUT("/api/v1/bank/transaction/:refid", handler.DisburseTransaction)
}

func (h *BankHandler) GetAccount(c echo.Context) error {
	accountno := c.Param("accountno")
	accountname := c.Param("accountname")
	ctx := c.Request().Context()

	account, err := h.bankService.GetAccountByNumberAndName(ctx, accountno, accountname)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, account)
}

func (h *BankHandler) CreateTransaction(c echo.Context) error {
	var transaction domain.Transaction
	err := c.Bind(&transaction)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	ctx := c.Request().Context()

	var ok bool
	if ok, err = isTransactionValid(&transaction); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.bankService.CreateTransaction(ctx, &transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, transaction)
}

func isTransactionValid(m *domain.Transaction) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *BankHandler) DisburseTransaction(c echo.Context) error {
	refID := c.Param("refid")
	ctx := c.Request().Context()

	err := h.bankService.DisburseTransaction(ctx, refID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}