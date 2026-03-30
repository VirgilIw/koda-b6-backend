package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

// CreateTransaction godoc
// @Summary Create new transaction
// @Description Create transaction with products
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateTransactionRequest true "Create Transaction Request"
// @Success 201 {object} dto.CreateTransactionResponse
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /admin/transactions [post]
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req dto.CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid request body",
			Error:   err.Error(),
		})
		return
	}

	userID := c.GetInt("user_id")

	code, err := h.service.CreateTransaction(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.CreateTransactionResponse{
		Success: true,
		Message: "transaction created successfully",
		Result: dto.TransactionResult{
			TransactionCode: code,
		},
	})
}
