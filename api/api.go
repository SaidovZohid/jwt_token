package api

import (
	"log"

	v1 "github.com/SaidovZohid/jwt_token/api/v1"
	"github.com/SaidovZohid/jwt_token/config"
	"github.com/SaidovZohid/jwt_token/storage"
	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RoutetOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

// @title           Swagger for User api
// @version         2.0
// @description     This is a user service api.
// @host      		localhost:8080
// @BasePath  		/v1
func New(opt *RoutetOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlersV1Options{
		Cfg:     opt.Cfg,
		Storage: &opt.Storage,
	})

	apiV1 := router.Group("/v1")
	log.Print("Hello")
	{
		apiV1.POST("/users", handlerV1.CreateUser)
		apiV1.GET("/users/:id", handlerV1.GetUser)

		apiV1.POST("/auth/register", handlerV1.Register)

		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler))
	}

	return router
}
