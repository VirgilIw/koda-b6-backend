package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
// 	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
// )

// type ForgotPwdHandler struct {
// 	service *service.ForgotPwdService
// }

// func NewForgotPwdHandler(service *service.ForgotPwdService) *ForgotPwdHandler {
// 	return &ForgotPwdHandler{
// 		service: service,
// 	}
// }

// // ForgotPassword godoc
// // @Summary Request OTP for forgot password
// // @Description Send OTP code to user's email for password reset
// // @Tags Authentication
// // @Accept json
// // @Produce json
// // @Param request body dto.ForgotPwdRequest true "Forgot Password Request"
// // @Success 200 {object} dto.ResponseForgotPwd
// // @Failure 400 {object} dto.ResponseForgotPwd
// // @Failure 500 {object} dto.ResponseForgotPwd
// // @Router /auth/forgot-password [post]
// func (f *ForgotPwdHandler) ForgotPassword(ctx *gin.Context) {

// 	var req dto.ForgotPwdRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, dto.ResponseForgotPwd{
// 			Success: false,
// 			Message: "invalid request body",
// 		})
// 		return
// 	}

// 	result, err := f.service.RequestForgotPassword(ctx.Request.Context(), req)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, dto.ResponseForgotPwd{
// 			Success: false,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, dto.ResponseForgotPwd{
// 		Success: true,
// 		Message: "OTP sent to email",
// 		Result:  result,
// 	})
// }

// // ResetPassword godoc
// // @Summary Reset user password
// // @Description Reset password using OTP sent to user's email
// // @Tags Authentication
// // @Accept json
// // @Produce json
// // @Param request body dto.ResetPasswordRequest true "Reset Password Request"
// // @Success 200 {object} dto.ResponseResetPwd "Password reset successfully"
// // @Failure 400 {object} dto.ResponseResetPwd "Invalid request or OTP"
// // @Failure 500 {object} dto.ResponseResetPwd "Internal server error"
// // @Router /auth/reset-password [patch]
// func (f *ForgotPwdHandler) ResetPassword(ctx *gin.Context) {

// 	var req dto.ResetPasswordRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, dto.ResponseResetPwd{
// 			Success: false,
// 			Message: "invalid request body",
// 		})
// 		return
// 	}

// 	err := f.service.ResetPassword(ctx.Request.Context(), req)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, dto.ResponseResetPwd{
// 			Success: false,
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, dto.ResponseResetPwd{
// 		Success: true,
// 		Message: "password reset successful",
// 	})
// }
