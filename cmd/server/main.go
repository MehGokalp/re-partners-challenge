package server

import (
	"fmt"
	_ "github.com/mehgokalp/re-partners-challenge/cmd/server/docs"
	"github.com/mehgokalp/re-partners-challenge/config"
	"github.com/mehgokalp/re-partners-challenge/internal/log"
	"github.com/mehgokalp/re-partners-challenge/internal/packaging/domain"
	"github.com/mehgokalp/re-partners-challenge/internal/server"
	"github.com/spf13/cobra"
)

// @title Messages API
// @version 1.0
// @description This is a sample server for managing orders.
// @host localhost:8081
// @BasePath /v1

func Server(
	cfg *config.Config,
	logger log.Logger,
	packagingHandler *domain.Handler,
) *cobra.Command {
	cmdName := "server"

	return &cobra.Command{
		Use:   cmdName,
		Short: "Run backend server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			r := server.GetRouter(
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
