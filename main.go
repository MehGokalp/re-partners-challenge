package main

import (
	"fmt"
	"github.com/mehgokalp/re-partners-challenge/cmd/server"
	"github.com/mehgokalp/re-partners-challenge/config"
	"github.com/mehgokalp/re-partners-challenge/internal/log"
	"github.com/mehgokalp/re-partners-challenge/internal/packaging/factory"
	"github.com/rotisserie/eris"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "re-partners-challenge",
		Short: "Main entry-point command for the application",
	}

	cfg := config.New()
	logger := log.New()

	packagingHandler := factory.NewDefaultHandler()

	rootCmd.AddCommand(
		server.Server(
			cfg,
			logger,
			&packagingHandler,
		),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(eris.ToString(err, true))
	}
}
