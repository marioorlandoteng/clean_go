package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	
	"moneytransfer/domain"
)

type BankService interface {
	GetAccountByNumberAndName(ctx context.Context, accountNumber string, accountName string) (domain.Account, error)
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) error
	UpdateTransaction(ctx context.Context, transaction *domain.Transaction) error
}

type BankHandler struct {
	bankService BankService
}

func NewBankHandler(e *echo.Echo, bankService BankService) {
	handler := &BankHandler{
		bankService: bankService,
	}
	e.GET("/account", handler.GetAccount)
	e.POST("/transaction", handler.CreateTransaction)
	e.PATCH("/transaction/:id", handler.UpdateTransaction)
}

func (h *BankHandler) GetAccount(c echo.Context) error {
	// todo
	return c.JSON(http.StatusOK, nil)
}

func (h *BankHandler) CreateTransaction(c echo.Context) error {
	// todo
	return c.NoContent(http.StatusNoContent)
}

func (h *BankHandler) UpdateTransaction(c echo.Context) error {
	// todo
	return c.NoContent(http.StatusNoContent)
}