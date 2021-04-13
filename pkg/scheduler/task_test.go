package scheduler

import (
	"context"
	"reflect"
	"testing"
)

func TestCreateTask(t *testing.T) {
	type args struct {
		ctx      context.Context
		name     string
		function TaskFunc
	}
	tests := []struct {
		name string
		args args
		want *ScheduledTask
	}{
		{
			name: "Simple Task",
			args: args{
				ctx:      context.TODO(),
				name:     "task_1",
				function: func(ctx context.Context) {},
			},
		},
		{
			name: "Task with panic",
			args: args{
				ctx:      context.TODO(),
				name:     "task_3",
				function: func(ctx context.Context) { panic("catch me!") },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTask(tt.args.ctx, tt.args.name, tt.args.function); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
