package main

import "testing"

func Test_isEscapePossible(t *testing.T) {
	type args struct {
		blocked [][]int
		source  []int
		target  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"simple pass",
			args:args{
				blocked: [][]int{
					[]int{0,1},
					[]int{0,2},
				},
				source:  []int{0,0},
				target:  []int{1,1},
			},
			want:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEscapePossible(tt.args.blocked, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("isEscapePossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
