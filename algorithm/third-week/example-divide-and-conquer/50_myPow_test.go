package example_divide_and_conquer

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "my pow",
			args: args{
				x: 2,
				n: 4,
			},
			want: 16,
		},
		{
			name: "my pow",
			args: args{
				x: 2,
				n: 1,
			},
			want: 2,
		},
		{
			name: "my pow",
			args: args{
				x: 2,
				n: -1,
			},
			want: 0.5,
		},
		{
			name: "my pow",
			args: args{
				x: 2,
				n: 0,
			},
			want: 1,
		},
		{
			name: "my pow",
			args: args{
				x: 2,
				n: 3,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
