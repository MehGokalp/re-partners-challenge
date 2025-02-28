package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mehgokalp/re-partners-challenge/cmd/server/docs"
	"github.com/mehgokalp/re-partners-challenge/config"
	pkgOrderPackaging "github.com/mehgokalp/re-partners-challenge/internal/delivery/http/order"
	"github.com/mehgokalp/re-partners-challenge/internal/log"
	"github.com/mehgokalp/re-partners-challenge/internal/meta"
	"github.com/mehgokalp/re-partners-challenge/internal/packaging"
	"github.com/spf13/cobra"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

// @title Messages API
// @version 1.0
// @description This is a sample server for managing orders.
// @host localhost:8081
// @BasePath /v1

func Server(
	cfg *config.Config,
	logger log.Logger,
	packagingHandler *packaging.Handler,
) *cobra.Command {
	cmdName := "server"

	return &cobra.Command{
		Use:   cmdName,
		Short: "Run backend server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			r := getRouter(
				logger,
				packagingHandler,
			)

			if err := r.Run(fmt.Sprintf(":%v", cfg.Port)); err != nil {
				return err
			}

			return nil
		},
	}
}

func getRouter(
	logger log.Logger,
	packagingHandler *packaging.Handler,
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

func jsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			line := make(map[string]interface{})

			line["app_name"] = meta.AppName
			line["level"] = "debug"
			line["status_code"] = params.StatusCode
			line["path"] = params.Path
			line["method"] = params.Method
			line["remote_addr"] = params.ClientIP
			line["response_time"] = params.Latency.String()
			line["time"] = params.TimeStamp.Format(time.RFC3339)

			s, _ := json.Marshal(line)
			return string(s) + "\n"
		},
	)
}
