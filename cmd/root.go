package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/thaitanloi365/go-logging"
	"github.com/thaitanloi365/go-monitor/config"
	"github.com/thaitanloi365/go-monitor/docker"
	"github.com/thaitanloi365/go-monitor/models"
	"github.com/thaitanloi365/go-monitor/scheduler"
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
	scheduler.New()

	rescheduleHealthCheckJobs()
}

func rescheduleHealthCheckJobs() {
	var jobs []*models.JobHealthCheck
	var err = models.GetDBInstance().Find(&jobs).Error
	if err != nil {
		logging.Global().Error(err)
		return
	}

	var scheduledJobs = scheduler.GetInstance().Jobs()
	for _, job := range jobs {
		if len(scheduledJobs) == 0 {
			j, err := scheduler.GetInstance().
				Every(5).
				Seconds().
				SetTag([]string{job.Tag}).
				Do(models.HeathCheckJobHandler, job.Endpoint, time.Duration(job.Timeout)*time.Second)
			if err != nil {
				logging.Global().Error(err)
				return
			}
			logging.Global().Info(j)
			logging.Global().Infof("Reschedule job tag = %s endpoint = %s interval = %d timeout = %d\n", job.Tag, job.Endpoint, job.Interval, job.Timeout)
			scheduledJobs = scheduler.GetInstance().Jobs()
		} else {
			for _, scheduledJob := range scheduledJobs {
				if scheduledJob.Tags()[0] == job.Tag {
					logging.Global().Infof("Skip Reschedule job tag = %s\n", job.Tag)
					continue
				}

				_, err := scheduler.GetInstance().
					Every(job.Interval).
					Seconds().
					SetTag([]string{job.Tag}).
					Do(models.HeathCheckJobHandler, job.Endpoint, time.Duration(job.Timeout)*time.Second)
				if err != nil {
					logging.Global().Error(err)
				}
				logging.Global().Infof("Reschedule job tag = %s endpoint = %s interval = %d timeout = %d\n", job.Tag, job.Endpoint, job.Interval, job.Timeout)
			}
		}

	}

}
