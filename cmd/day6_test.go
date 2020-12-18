package cmd

import (
	"testing"
)

func TestDay6p1(t *testing.T) {
	d := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
	x := Day6Part1(d)
	t.Logf("Result: %v\n", x)
	if x != 11 {
		t.FailNow()
	}
}

func TestDay6p2(t *testing.T) {
	d := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
	x := Day6Part2(d)
	t.Logf("Result: %v\n", x)
	if x != 6 {
		t.FailNow()
	}
}

func TestDay6Part1(t *testing.T) {
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
			if got := Day6Part1(tt.args.input); got != tt.want {
				t.Errorf("Day6Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay6Part2(t *testing.T) {
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
			if got := Day6Part2(tt.args.input); got != tt.want {
				t.Errorf("Day6Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
