package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var _serveCMD = &cobra.Command{
	Use:   "serve",
	Short: "serve APIs",
	RunE: func(cmd *cobra.Command, args []string) error {
		serve()
		return nil
	},
}

func serve() {
	fmt.Println("serving a web server")
}
