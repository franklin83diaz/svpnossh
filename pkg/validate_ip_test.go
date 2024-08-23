package pkg

import "testing"

func TestValidateIp(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{name: "Test 1", args: args{ip: "200.55.0.2"}, want: true},
		{name: "Test 2", args: args{ip: "192.168.0.2"}, want: true},
		{name: "Test 3", args: args{ip: "172.16.0.2"}, want: true},
		{name: "Test 4", args: args{ip: "127.0.0.1"}, want: true},
		{name: "Test 5", args: args{ip: "0.0.0.0"}, want: false},
		{name: "Test 6", args: args{ip: ""}, want: false},
		{name: "Test 7", args: args{ip: "This is not an IP"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateIp(tt.args.ip); got != tt.want {
				t.Errorf("ValidateIp() = %v, want %v", got, tt.want)
			}
		})
	}
}
