package module01

import (
	"fmt"
	"testing"
)

func TestTape(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want int
	}{
		{"All negatives", []int{-5, -3, -1, -4}, 3},
		{"Zeroes", []int{0, 0}, 0},
		{"Pos/neg", []int{-2, 2}, 4},
		{"Small values", []int{0, 1, 0, 1, 0, 1, 99}, 96},
		{"Provided example", []int{3, 1, 2, 4, 3}, 1},
		{"Provided negatives", []int{-3, -1, -2, -4, -3}, 1},
		{"Small elements", []int{-10, -20, -30, -40, 100}, 20},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Tape: %s\n", tc.name), func(t *testing.T) {
			got := Tape(tc.A)
			if got != tc.want {
				t.Fatalf("Tape%s | (%v) = %v; want %v", tc.name, tc.A, got, tc.want)
			}
		})
	}
}
