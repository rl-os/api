package jobs

import (
	"sync"
	"time"
)

const (
	JOB_STATUS_IS_PENDING     = "pending"
	JOB_STATUS_IS_IN_PROGRESS = "in_progress"
	JOB_STATUS_IS_SUCCESS     = "success"
	JOB_STATUS_IS_ERROR       = "error"
	JOB_STATUS_IS_CANCELED    = "canceled"
)

type Job struct {
	Id       string            `json:"id"`
	Type     string            `json:"type"`
	CreateAt time.Time         `json:"create_at"`
	StartAt  time.Time         `json:"start_at"`
	Data     map[string]string `json:"data"`
}

type Worker interface {
	Run()
	Stop()
	JobChannel() chan<- Job
}

type Workers struct {
	startOnce sync.Once
}
