package handler

import (
	"net/http"

	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.transactionService.CreateTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil menambahkan transaksi"})
}

func (h *TransactionHandler) HistoryTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	history, err := h.transactionService.HistoryTransaction(transaction.NIK)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, history)
}
