package handler

import (
	"net/http"

	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

type DepositHandler struct {
	depoService service.DepositService
}

func NewDepositHandler(depoService service.DepositService) *DepositHandler {
	return &DepositHandler{depoService}
}

func (h *DepositHandler) Create(c *gin.Context) {
	var depo model.Deposit
	if err := c.ShouldBindJSON(&depo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.depoService.CreateDeposit(&depo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deposit created successfully"})
}

func (h *DepositHandler) GetDepositByNik(c *gin.Context) {
	var depo model.Deposit
	if err := c.ShouldBindJSON(&depo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deposit, err := h.depoService.GetDepositByNik(depo.NIK)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deposit)
}

func (h *DepositHandler) Update(c *gin.Context) {
	var depo model.Deposit
	if err := c.ShouldBindJSON(&depo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.depoService.UpdateDeposit(&depo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, depo)
}

func (h *DepositHandler) Delete(c *gin.Context) {
	var depo model.Deposit
	if err := c.ShouldBindJSON(&depo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.depoService.DeleteDeposit(depo.NIK); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deposit deleted successfully"})
}
