package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type AuthHandler struct {
	authService   *service.AuthService
	forgotService *service.ForgotPwdService
}

func NewAuthHandler(authService *service.AuthService, forgotService *service.ForgotPwdService) *AuthHandler {
	return &AuthHandler{
		authService:   authService,
		forgotService: forgotService,
	}
}

// Auth login godoc
// @Summary      Login
// @Description  Authenticate user and return JWT token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request  body      dto.AuthLoginRequest  true  "Login request"
// @Success      200      {object}  dto.ResponseToken
// @Failure      400      {object}  dto.ResponseToken
// @Failure      401      {object}  dto.ResponseToken
// @Router       /auth/login [post]
func (l *AuthHandler) AuthLogin(ctx *gin.Context) {
	var data dto.AuthLoginRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseToken{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	token := l.authService.AuthLogin(ctx, data.Email, data.Password)

	if token != "" {
		ctx.JSON(http.StatusOK, dto.ResponseToken{
			Success: true,
			Message: "login success",
			Token:   token,
		})
		return
	}

	ctx.JSON(http.StatusUnauthorized, dto.ResponseToken{
		Success: false,
		Message: "login failed",
		Error:   "invalid email or password",
	})
}

// ForgotPassword godoc
// @Summary Request OTP for forgot password
// @Description Send OTP code to user's email for password reset
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.ForgotPwdRequest true "Forgot Password Request"
// @Success 200 {object} dto.ResponseForgotPwd
// @Failure 400 {object} dto.ResponseForgotPwd
// @Router /auth/forgot-password [post]
func (a *AuthHandler) ForgotPassword(ctx *gin.Context) {

	var req dto.ForgotPwdRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseForgotPwd{
			Success: false,
			Message: "invalid request body",
		})
		return
	}

	result, err := a.forgotService.RequestForgotPassword(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseForgotPwd{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseForgotPwd{
		Success: true,
		Message: "OTP sent to email",
		Result:  result,
	})
}

// ResetPassword godoc
// @Summary Reset user password
// @Description Reset password using OTP sent to user's email
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.ResetPasswordRequest true "Reset Password Request"
// @Success 200 {object} dto.ResponseResetPwd
// @Failure 400 {object} dto.ResponseResetPwd
// @Router /auth/reset-password [patch]
func (a *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req dto.ResetPasswordRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseResetPwd{
			Success: false,
			Message: "invalid request body",
		})
		return
	}

	err := a.forgotService.ResetPassword(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseResetPwd{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseResetPwd{
		Success: true,
		Message: "password reset successful",
	})
}
