package handler

import (
	"net/http"
	"strconv"

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

// GetTransactionsByUserID godoc
// @Summary Get user transactions
// @Description Get all transactions for authenticated user
// @Tags Transactions
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.GetTransactionResponse
// @Failure 401 {object} dto.Response "Unauthorized"
// @Failure 500 {object} dto.Response "Internal server error"
// @Router /transactions [get]
func (h *TransactionHandler) GetTransactionsByUserID(c *gin.Context) {

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "invalid user",
		})
		return
	}

	result, err := h.service.GetTransactionsByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "failed to get transactions",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GetTransactionResponse{
		Success: true,
		Message: "success",
		Result:  result,
	})
}

// GetTransactionDetail godoc
// @Summary Get transaction detail
// @Description Get transaction detail by ID (user only)
// @Tags Transactions
// @Security BearerAuth
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} dto.GetTransactionDetailResponse
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Transaction not found"
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransactionDetail(c *gin.Context) {

	// 1. ambil user_id dari JWT
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "unauthorized",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "invalid user",
		})
		return
	}

	// 2. ambil param id
	idParam := c.Param("id")
	transactionID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid transaction id",
		})
		return
	}

	// 3. call service
	result, err := h.service.GetTransactionDetail(c.Request.Context(), userID, transactionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "transaction not found",
		})
		return
	}

	// 4. response
	c.JSON(http.StatusOK, dto.GetTransactionDetailResponse{
		Success: true,
		Message: "success",
		Result:  *result,
	})
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
// @Router /transactions [post]
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

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "invalid user",
		})
		return
	}

	tId, code, err := h.service.CreateTransaction(c.Request.Context(), userID, req)
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
		Result: dto.CreateTransactionResult{
			ID:              tId,
			TransactionCode: code,
		},
	})
}

// UpdateTransactionStatus godoc
// @Summary Update transaction status
// @Description Update status of transaction (admin)
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param request body dto.UpdateTransactionStatusRequest true "Update Status Request"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /admin/transactions/{id}/status [patch]
func (h *TransactionHandler) UpdateTransactionStatus(c *gin.Context) {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid transaction id",
		})
		return
	}

	var req dto.UpdateTransactionStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "invalid request body",
			Error:   err.Error(),
		})
		return
	}

	if req.Status == "" || (req.Status != "done" && req.Status != "on progress") {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "status must be either 'done' or 'on progress'",
		})
		return
	}

	if err := h.service.UpdateTransactionStatus(c.Request.Context(), transactionID, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "transaction status updated successfully",
	})
}
