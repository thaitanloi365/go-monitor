package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/thaitanloi365/go-monitor/helper"
)

// Scheduler new scheduler
type Scheduler struct {
	*gocron.Scheduler
}

var instance *Scheduler

// New init
func New() {
	tz, err := time.LoadLocation(helper.DefaultTimezone)
	if err != nil {
		fmt.Println("Load timezone error: %+v", err)
		tz = time.UTC
	}
	var s = gocron.NewScheduler(tz)
	instance = &Scheduler{
		s,
	}

	s.StartAsync()
}

// GetInstance get instance
func GetInstance() *Scheduler {
	if instance == nil {
		panic("Must be call New() first")
	}

	return instance
}
