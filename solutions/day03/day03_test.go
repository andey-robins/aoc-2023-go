package day03

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		fname string
		want  int
	}{
		{"../../inputs/tests/03.txt", 4361},
		{"../../inputs/03.txt", 520135},
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
		{"../../inputs/tests/03.txt", 467835},
		{"../../inputs/03.txt", 72514855},
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
