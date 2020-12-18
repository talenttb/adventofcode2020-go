package cmd

import (
	"testing"
)

func TestDay3p1(t *testing.T) {
	d := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	x := Day3Part1(d)
	t.Logf("Result: %v\n", x)
}

func TestDay3p2(t *testing.T) {
	d := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	x := Day3Part2(d)
	t.Logf("Result: %v\n", x)
}

func TestDay3Part1(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day3Part1(tt.args.arr); got != tt.want {
				t.Errorf("Day3Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay3Part2(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day3Part2(tt.args.arr); got != tt.want {
				t.Errorf("Day3Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day3Part2(t *testing.T) {
	type args struct {
		arr []string
		s   steps
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day3Part2(tt.args.arr, tt.args.s); got != tt.want {
				t.Errorf("day3Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
