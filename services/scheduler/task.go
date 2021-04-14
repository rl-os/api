package scheduler

import (
	"context"
	"fmt"
	"time"
)

type TaskFunc func(ctx context.Context)

type ScheduledTask struct {
	Name     string        `json:"name"`
	Delay    time.Duration `json:"delay"`
	Interval time.Duration `json:"interval"`

	function  TaskFunc
	context   context.Context
	cancel    chan struct{}
	cancelled chan struct{}
}

func (task ScheduledTask) WithDelay(delay time.Duration) *ScheduledTask {
	task.Delay = delay
	return &task
}

func (task ScheduledTask) WithInterval(duration time.Duration) *ScheduledTask {
	task.Interval = duration

	return &task
}

// CreateTask with default params (interval 5s and without delay)
func CreateTask(ctx context.Context, name string, function TaskFunc) *ScheduledTask {
	return &ScheduledTask{
		Name:      name,
		Delay:     time.Nanosecond,
		Interval:  time.Second * 5,
		function:  function,
		context:   ctx,
		cancel:    make(chan struct{}),
		cancelled: make(chan struct{}),
	}
}

func (task *ScheduledTask) Cancel() {
	task.cancel <- struct{}{}
	<-task.cancelled
}

func (task *ScheduledTask) String() string {
	return fmt.Sprintf(
		"%s\nInterval: %s\nDelay: %s\n",
		task.Name,
		task.Interval.String(),
		task.Delay.String(),
	)
}

func (task ScheduledTask) Run() {
	ctx, cancelCtx := context.WithCancel(task.context)

	go func() {
		defer close(task.cancelled)

		ticker := time.NewTicker(task.Interval)
		defer func() {
			ticker.Stop()
			cancelCtx()
			task.cancelled <- struct{}{}
		}()

		// task delay
		time.Sleep(task.Delay)

		// first call
		task.function(ctx)

	loop:
		for {
			select {
			case <-ticker.C:
				task.function(ctx)
			case <-ctx.Done():
			case <-task.cancel:
				task.cancelled <- struct{}{}
				break loop
			}
		}
	}()
}
