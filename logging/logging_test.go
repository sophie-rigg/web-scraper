package logging

import (
	"testing"
)

func TestLogLevel_Set(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		l       LogLevel
		args    args
		wantErr bool
	}{
		{
			name: "Valid log level",
			l:    LogLevel("info"),
			args: args{
				s: "info",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.Set(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.l.Value() != tt.args.s {
				t.Errorf("Set() got = %v, want %v", tt.l.String(), tt.args.s)
			}
		})
	}
}
