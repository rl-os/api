package utils

import (
	"reflect"
	"testing"
)

func TestIfThen(t *testing.T) {
	type args struct {
		condition bool
		a         interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "if true return 1", args: args{condition: true, a: 1}, want: 1},
		{name: "if true return 0", args: args{condition: true, a: 0}, want: 0},
		{name: "if false return 1", args: args{condition: false, a: 1}, want: nil},
		{name: "if false return 0", args: args{condition: false, a: 0}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfThen(tt.args.condition, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfThen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfThenElse(t *testing.T) {
	type args struct {
		condition bool
		a         interface{}
		b         interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "if true return 1 else 0", args: args{condition: true, a: 1}, want: 1},
		{name: "if false return 1 else 0", args: args{condition: true, a: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfThenElse(tt.args.condition, tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfThenElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultIfNil(t *testing.T) {
	type args struct {
		value        interface{}
		defaultValue interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "if nil return 1", args: args{value: nil, defaultValue: 1}, want: 1},
		{name: "if not nil return 1", args: args{value: true, defaultValue: 1}, want: true},
		{name: "if nil return something", args: args{value: nil, defaultValue: "something"}, want: "something"},
		{name: "if not nil return something", args: args{value: true, defaultValue: "something"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultIfNil(tt.args.value, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultIfNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstNonNil(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "first not nil", args: args{values: []interface{}{"1"}}, want: "1"},
		{name: "first not nil", args: args{values: []interface{}{nil, nil, "1"}}, want: "1"},
		{name: "first not nil", args: args{values: []interface{}{nil, nil, "1", "2"}}, want: "1"},
		{name: "first not nil", args: args{values: []interface{}{nil, nil, "1", nil, "2"}}, want: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstNonNil(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstNonNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
