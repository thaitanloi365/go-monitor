package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thaitanloi365/go-monitor/routes"
	"github.com/thaitanloi365/go-monitor/sse"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		sse.New()
		routes.SetupRoutes()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
