package internal

import (
	"sync"
	"time"
)

type Crontab struct {

	//crontab file check
	monitorModifyTicker *time.Ticker

	ticker *time.Ticker

	task []task

	lock sync.Mutex
}

func New() {

}

func (c *Crontab) Run() {

}

func (c *Crontab) ShutDown() {

}

func (c *Crontab) Start() {

}

func (c *Crontab) AddTask(t task) error {

	return nil
}

// 周期性监控crontab的变动
func (c *Crontab) PeriodicMonitorCrontab() {

}
