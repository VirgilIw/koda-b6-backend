package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type ForgotPwdHandler struct {
	service *service.ForgotPwdService
}

func NewForgotPwdHandler(service *service.ForgotPwdService) *ForgotPwdHandler {
	return &ForgotPwdHandler{
		service: service,
	}
}

// ForgotPassword godoc
// @Summary      Request OTP for forgot password
// @Description  Endpoint to request a one-time password (OTP) when a user forgets their password. The OTP will be sent to the user's email.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request  body     dto.ForgotPwdRequest  true  "Forgot Password Request"
// @Success      200     {object}  dto.ResponseForgotPwd "OTP sent successfully (do not expose OTP)"
// @Failure      400     {object}  dto.ResponseForgotPwd   "Bad request or email not found"
// @Failure      500     {object}  dto.ResponseForgotPwd   "Internal server error"
// @Router       /auth/forgot-password [post]
func (f *ForgotPwdHandler) ForgotPassword(ctx *gin.Context) {

	var req dto.ForgotPwdRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseForgotPwd{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	fpwd, err := f.service.RequestForgotPassword(ctx.Request.Context(), req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseForgotPwd{
			Success: false,
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseForgotPwd{
		Success: true,
		Message: "success forgot password",
		Result:  fpwd,
	})
}
