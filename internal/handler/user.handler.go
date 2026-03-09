package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// GetUsers godoc
// @Summary      Get users
// @Description  Get users
// @Tags         Users
// @Produce      json
// @Success      200  {object}  dto.Response
// @Failure      400  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Security     BearerAuth
// @Router       /users [get]
func (u *UserHandler) GetUsers(ctx *gin.Context) {
	datas, err := u.service.GetUsers(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "success get users",
		Result:  datas,
	})
}
