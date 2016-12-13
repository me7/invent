package main

import "testing"

func Test_must(t *testing.T)                  {}
func Test_init(t *testing.T)                  {}
func Test_main(t *testing.T)                  {}
func Test_mainHandler_ServeHTTP(t *testing.T) {}

func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1+2", args{1, 2}, 3},
		{"2+2", args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
