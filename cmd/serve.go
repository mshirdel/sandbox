package cmd

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mshirdel/sandbox/app"
	api "github.com/mshirdel/sandbox/app/http"
	"github.com/spf13/cobra"
)

var _serveCMD = &cobra.Command{
	Use:   "serve",
	Short: "serve APIs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve()
	},
}

func serve() error {
	app := app.New()
	server := api.NewHTTPServer(app)

	defer server.Shutdown()
	go server.Start()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-ctx.Done()
	return nil
}
