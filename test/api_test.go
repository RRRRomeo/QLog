package api

import (
	"sync"
	"testing"

	"github.com/QLog/internal"
)

func TestLogApi_Debugf(t *testing.T) {
	type fields struct {
		Mutex  sync.Mutex
		logger *internal.Logger
	}
	type args struct {
		fmt string
		v   []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := &LogApi{
				Mutex:  tt.fields.Mutex,
				logger: tt.fields.logger,
			}
			log.Debugf(tt.args.fmt, tt.args.v...)
		})
	}
}

func TestDebugf(t *testing.T) {
	type args struct {
		fmt string
		v   []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugf(tt.args.fmt, tt.args.v...)
		})
	}
}
