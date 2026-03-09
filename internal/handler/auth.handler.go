package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
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

	token := l.service.AuthLogin(ctx, data.Email, data.Password)

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
