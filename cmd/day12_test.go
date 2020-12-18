package cmd

import (
	"testing"
)

func TestDay12p1(t *testing.T) {
	d := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	x := Day12Part1(d)
	t.Logf("Result: %v\n", x)
	if x != 25 {
		t.Fail()
	}
}

func TestDay12p2(t *testing.T) {
	d := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	x := Day12Part2(d)
	t.Logf("Result: %v\n", x)
	if x != 286 {
		t.Fail()
	}
}

func TestDay12p2v2(t *testing.T) {
	d := []string{
		"F42",
		"L180",
		"S4",
		"W4",
		"N5",
		"R270",
		"F10",
	}
	x := Day12Part2(d)
	t.Logf("Result: %v\n", x)
	if x != 498 {
		t.Fail()
	}
}

func TestDay12Part1(t *testing.T) {
	type args struct {
		input []string
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
			if got := Day12Part1(tt.args.input); got != tt.want {
				t.Errorf("Day12Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay12Part2(t *testing.T) {
	type args struct {
		input []string
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
			if got := Day12Part2(tt.args.input); got != tt.want {
				t.Errorf("Day12Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
