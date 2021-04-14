package scheduler

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {
	type args struct {
		ctx      context.Context
		name     string
		function TaskFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple task",
			args: args{
				ctx:      context.TODO(),
				name:     "task_1",
				function: func(ctx context.Context) {},
			},
		},
		{
			name: "task with panic",
			args: args{
				ctx:      context.TODO(),
				name:     "task_3",
				function: func(ctx context.Context) { panic("catch me!") },
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(tt.args.ctx)

			var mx sync.Mutex
			val := int64(0)

			ff := tt.args.function
			task := CreateTask(ctx, tt.args.name, func(ctx context.Context) {
				ff(ctx)

				mx.Lock()
				val++
				mx.Unlock()
			}).
				WithInterval(time.Millisecond)

			defer func() {
				cancel()
			}()

			task.Run()

			<-time.Tick(time.Millisecond * 10)

			mx.Lock()
			result := val
			mx.Unlock()

			if result == 0 && !tt.wantErr {
				t.Errorf("invalid interval")
			}
		})
	}
}
