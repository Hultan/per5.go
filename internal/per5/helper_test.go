package per5

import (
	"testing"
)

func TestPer5_Map(t *testing.T) {
	type args struct {
		v            float64
		iMin         float64
		iMax         float64
		oMin         float64
		oMax         float64
		withinBounds []bool
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"map(2,0,4,10,20) == 15", args{2, 0, 4, 10, 20, []bool{}}, 15},
		{"map(-1,0,4,10,20) == 5", args{-1, 0, 4, 10, 20, []bool{}}, 7.5},
		{"map(5,0,4,10,20) == 22.5", args{5, 0, 4, 10, 20, []bool{}}, 22.5},
		{"map(2,0,4,10,20,false) == 15", args{2, 0, 4, 10, 20, []bool{false}}, 15},
		{"map(-1,0,4,10,20,false) == 5", args{-1, 0, 4, 10, 20, []bool{false}}, 7.5},
		{"map(5,0,4,10,20,false) == 22.5", args{5, 0, 4, 10, 20, []bool{false}}, 22.5},
		{"map(2,0,4,10,20,true) == 15", args{2, 0, 4, 10, 20, []bool{true}}, 15},
		{"map(-1,0,4,10,20,true) == 10", args{-1, 0, 4, 10, 20, []bool{true}}, 10},
		{"map(5,0,4,10,20,true) == 20", args{5, 0, 4, 10, 20, []bool{true}}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Per5{}
			if got := p.Map(tt.args.v, tt.args.iMin, tt.args.iMax, tt.args.oMin, tt.args.oMax, tt.args.withinBounds...); got != tt.want {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPer5_Constraint(t *testing.T) {
	type args struct {
		v   float64
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Constraint(5,0,10) == 5", args{5, 0, 10}, 5},
		{"Constraint(-5,0,10) == 0", args{-5, 0, 10}, 0},
		{"Constraint(15,0,10) == 10", args{15, 0, 10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Per5{}
			if got := p.Constraint(tt.args.v, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Constraint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPer5_Random(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
		min  float64
		max  float64
	}{
		{"1,2", args{1, 2}, 1, 2},
		{"10,20", args{10, 20}, 10, 20},
		{"-10,10", args{-10, 10}, -10, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Per5{}
			for i := 0; i < 100; i++ {
				if got := p.Random(tt.args.min, tt.args.max); got < tt.min || got > tt.max {
					t.Errorf("Random() = %v, want %v <= %v <= %v", got, tt.min, got, tt.max)
				}
			}
		})
	}
}
