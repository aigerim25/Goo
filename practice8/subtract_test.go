package main

import "testing"

func TestSubtractTableDriven(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"both positive", 4, 2, 2},
		{"both negative", -4, -2, -2},
		{"negative + positive", -4, 2, -6},
		{"positive + negative", 4, -2, 6},
		{"negative minus 0", -4, 0, -4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}

}
