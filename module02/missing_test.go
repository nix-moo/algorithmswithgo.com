package module02

import (
	"fmt"
	"testing"
)

func TestMissing(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want int
	}{
		{"One element", []int{2}, 1},
		{"Missing first element", []int{2, 3, 5, 4}, 1},
		{"Missing last element", []int{2, 3, 1, 5, 1}, 6},
		{"Empty array", []int{}, 1},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintln("Missing element"), func(t *testing.T) {
			got := Missing(tc.A)
			if got != tc.want {
				t.Fatalf("Missing(%v) = %v; want %v", tc.A, got, tc.want)
			}
		})
	}
}
