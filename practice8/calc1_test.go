// table-driven test
package main

import "testing"

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"both positive", 2, 3, 5},
		{"positive + zero", 5, 0, 5},
		{"negative + positive", -1, 4, 3},
		{"both negative", -2, -3, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d,%d) = %d, want %d ", tt.a, tt.b, got, tt.want) // %d говорит нам о том, чтобы мы вставили сюда число, целое число
			}
		})
	}
}
