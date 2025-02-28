package server

import (
	"github.com/gin-gonic/gin"
	pkgOrderPackaging "github.com/mehgokalp/re-partners-challenge/internal/delivery/http/order"
	"github.com/mehgokalp/re-partners-challenge/internal/log"
	"github.com/mehgokalp/re-partners-challenge/internal/packaging/domain"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRouter(
	logger log.Logger,
	packagingHandler *domain.Handler,
) *gin.Engine {
	r := gin.New()
	r.Use(gin.ErrorLogger())
	r.Use(jsonLoggerMiddleware())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")

	v1.GET("/calculate-packaging/", pkgOrderPackaging.NewHandler(logger, packagingHandler))

	return r
}
