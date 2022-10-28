package dp

import "testing"

func TestBag01(t *testing.T) {
	type args struct {
		weight []int
		value  []int
		bag    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: args{
			weight: []int{1, 3, 4, 5},
			value:  []int{15, 20, 30, 55},
			bag:    6,
		},
			want: 70,
		},

		{name: "t2", args: args{
			weight: []int{1, 3, 4, 5},
			value:  []int{15, 20, 30, 20},
			bag:    6,
		},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bag01(tt.args.weight, tt.args.value, tt.args.bag); got != tt.want {
				t.Errorf("Bag01() = %v, want %v", got, tt.want)
			}
		})
	}
}
