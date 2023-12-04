package day04

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		fname string
		want  int
	}{
		{"../../inputs/tests/04.txt", 13},
		{"../../inputs/04.txt", 24542},
	}
	for _, tt := range tests {
		t.Run(tt.fname, func(t *testing.T) {
			input, err := os.ReadFile(tt.fname)
			if err != nil {
				t.Errorf("Part1() error = %v", err)
			}

			if got := Part1(string(input)); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		fname string
		want  int
	}{
		{"../../inputs/tests/04.txt", 30},
		{"../../inputs/04.txt", 8736438},
	}
	for _, tt := range tests {
		t.Run(tt.fname, func(t *testing.T) {
			input, err := os.ReadFile(tt.fname)
			if err != nil {
				t.Errorf("Part2() error = %v", err)
			}

			if got := Part2(string(input)); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
