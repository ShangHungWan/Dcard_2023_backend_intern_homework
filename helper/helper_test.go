package helper

import (
	"net/http"
	"testing"
)

func TestIsStatusOK(t *testing.T) {
	type args struct {
		status int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				status: http.StatusOK,
			},
			want: true,
		},
		{
			name: "created",
			args: args{
				status: http.StatusCreated,
			},
			want: true,
		},
		{
			name: "no content",
			args: args{
				status: http.StatusNoContent,
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				status: http.StatusNotFound,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStatusOK(tt.args.status); got != tt.want {
				t.Errorf("IsStatusOK() = %v, want %v", got, tt.want)
			}
		})
	}
}
