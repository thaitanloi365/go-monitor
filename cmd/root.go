package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thaitanloi365/go-logging"
	"github.com/thaitanloi365/go-monitor/config"
	"github.com/thaitanloi365/go-monitor/docker"
	"github.com/thaitanloi365/go-monitor/models"
)

// RootCmd root command
var RootCmd = &cobra.Command{
	Use:   "Go-monitor",
	Short: "Go-monitor",
}

var cfgFile string
var initOnStartup bool
var pwd string

// Execute Execute root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pwd = dir

	cobra.OnInitialize(setup)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is %s/%s.yaml)", dir, "config"))
	RootCmd.PersistentFlags().BoolVar(&initOnStartup, "initOnStartup", true, "Init all dependency on startup")
}

func setup() {
	logging.NewGlobal()
	config.New(cfgFile)
	docker.New()
	models.Setup()
}
