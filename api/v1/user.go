package v1

import (
	"net/http"
	"strconv"

	"github.com/SaidovZohid/jwt_token/api/models"
	"github.com/SaidovZohid/jwt_token/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /users [post]
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.GetUser
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateUser(ctx *gin.Context) {
	var (
		req models.User
	)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})	
		return 
	}

	user, err := h.Storage.User().Create(&repo.User{
		Name: req.Name,
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})	
		return 
	}

	ctx.JSON(http.StatusOK, models.GetUser{
		Id: user.ID,
		Name: req.Name,
		Username: req.Name,
		Email: req.Email,
		CreatedAt: user.CreatedAt,
	})
}

// @Router /users/{id} [get]
// @Summary Get a user by it's id
// @Description Get a user by it's id
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {object} models.GetUser
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})	
		return 
	}

	user, err := h.Storage.User().Get(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})	
		return 
	}

	ctx.JSON(http.StatusOK, models.GetUser{
		Id: user.ID,
		Name: user.Name,
		Username: user.Username,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	})
}
