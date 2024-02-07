package utils

import (
	"testing"
)

func TestCheckDomainMatch(t *testing.T) {
	type args struct {
		domain string
		url    string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Valid https url",
			args: args{
				domain: "https://www.monzo.com",
				url:    "https://www.monzo.com/legal",
			},
			want:  "https://www.monzo.com/legal",
			want1: true,
		},
		{
			name: "Valid non https url",
			args: args{
				domain: "https://www.monzo.com",
				url:    "/legal/",
			},
			want:  "https://www.monzo.com/legal/",
			want1: true,
		},
		{
			name: "Invalid url",
			args: args{
				domain: "https://www.monzo.com",
				url:    "https://www.google.com",
			},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckDomainMatch(tt.args.domain, tt.args.url)
			if got != tt.want {
				t.Errorf("CheckDomainMatch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckDomainMatch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
