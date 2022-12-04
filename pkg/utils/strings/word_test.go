package strings

import "testing"

func TestCapitalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test Capitalize string 1", args: args{s: "google"}, want: "Google"},
		{name: "Test Capitalize string 2", args: args{s: "profiler By us"}, want: "Profiler By us"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.args.s); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
