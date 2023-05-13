package api

import (
	"testing"

	"github.com/QLog/api"
)

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
			api.Debugf(tt.args.fmt, tt.args.v...)
		})
	}
}
