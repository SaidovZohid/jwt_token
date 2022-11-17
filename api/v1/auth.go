package v1

import (
	"net/http"
	"time"

	"github.com/SaidovZohid/jwt_token/api/models"
	"github.com/SaidovZohid/jwt_token/pkg/utils"
	"github.com/SaidovZohid/jwt_token/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /auth/register [post]
// @Summary Register a user
// @Description Register a user
// @Tags register
// @Accept json
// @Produce json
// @Param user body models.RegisterUser true "User"
// @Success 200 {object} models.RegisterResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Register(ctx *gin.Context) {
	var (
		req models.RegisterUser
	)

	err := ctx.ShouldBindJSON(&req)

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	result, err := h.Storage.User().Create(&repo.User{
		Name: req.Name,
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	token, _, err := utils.CreateToken(req.Username, req.Email, 30 *time.Second)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.RegisterResponse{
		ID: result.ID,
		Name: req.Name,
		Email: req.Email,
		Username: req.Username,
		CreatedAt: result.CreatedAt,
		AccessToken: token,
	})
}
