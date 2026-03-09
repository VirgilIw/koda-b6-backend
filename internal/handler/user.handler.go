package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

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

// UpdateProfile godoc
// @Summary      Update profile
// @Description  Update the logged-in user's profile, including optional picture upload
// @Tags         Users
// @Accept       multipart/form-data
// @Produce      json
// @Param        fullname   formData  string  false  "Full name"
// @Param        email      formData  string  false  "Email"
// @Param        password   formData  string  false  "Password"
// @Param        phone      formData  string  false  "Phone number"
// @Param        address    formData  string  false  "Address"
// @Param        role       formData  string  false  "Role"
// @Param        picture    formData  file    false  "Profile picture"
// @Success      200        {object}  dto.ResponseOneData
// @Failure      400        {object}  dto.ResponseOneData
// @Failure      500        {object}  dto.ResponseOneData
// @Security     BearerAuth
// @Router       /users [patch]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	// Ambil userID dari JWT
	idRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}
	userId := idRaw.(int)

	// Ambil form data
	var req dto.UpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request: " + err.Error(),
		})
		return
	}
	req.Id = userId

	file, err := c.FormFile("picture")
	if err == nil {

		// agar nama file tidak sama gunaakn time.now.unix
		// filepath.base = menghapus bagian folder/path, menyisakan nama file terakhir serta ekstensi filenya.
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		path := "./images/"
		if err := c.SaveUploadedFile(file, path+filename); err != nil {
			c.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Message: "failed to upload picture",
			})
			return
		}
		req.Picture = &path
	}

	// Update via service
	updatedUser, err := h.service.UpdateProfile(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "failed to update profile",
			Error:   err.Error(),
		})
		return
	}

	// **Konversi model.UserModel ke dto.Users**
	respUser := dto.Users{
		Id:       updatedUser.Id,
		FullName: updatedUser.FullName,
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
		Picture:  updatedUser.Picture,
		Phone:    updatedUser.Phone,
		Address:  updatedUser.Address,
		Role:     updatedUser.Role,
	}

	c.JSON(http.StatusOK, dto.ResponseOneData{
		Success: true,
		Message: "profile updated",
		Result:  respUser,
	})
}
