package main

import "testing"

func TestDivideTableDriven(t *testing.T) {
	tests := []struct {
		name  string
		a, b  int
		want  int
		Error bool
	}{
		{"both positive", 10, 5, 2, false},
		{"both negative", -20, -5, 4, false},
		{"negative and positive", -30, 6, -5, false},
		{"positive with 0", 0, 5, 0, false},
		{"division by 0", 10, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if tt.Error {
				if err == nil {
					t.Errorf("expected an error")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Divide(%d,%d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
