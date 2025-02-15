package main

import (
	"testing"
)

func TestIsIPInList(t *testing.T) {
	tests := []struct {
		name     string
		clientIP string
		ipList   []string
		want     bool
	}{
		{
			name:     "Single IP match",
			clientIP: "192.168.1.1",
			ipList:   []string{"192.168.1.1"},
			want:     true,
		},
		{
			name:     "CIDR match",
			clientIP: "192.168.1.100",
			ipList:   []string{"192.168.1.0/24"},
			want:     true,
		},
		{
			name:     "No match",
			clientIP: "10.0.0.1",
			ipList:   []string{"192.168.1.0/24", "172.16.0.1"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIPInList(tt.clientIP, tt.ipList); got != tt.want {
				t.Errorf("isIPInList() = %v, want %v", got, tt.want)
			}
		})
	}
}
