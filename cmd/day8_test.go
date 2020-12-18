package cmd

import (
	"testing"
)

func TestDay8p1(t *testing.T) {
	d := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
	x := Day8Part1(d)
	t.Logf("Result: %v\n", x)
	if x != 5 {
		t.Fail()
	}
}

func TestDay8p2(t *testing.T) {
	d := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
	x := Day8Part2(d)
	t.Logf("Result: %v\n", x)
	if x != 8 {
		t.Fail()
	}
}

func TestDay8Part1(t *testing.T) {
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
			if got := Day8Part1(tt.args.input); got != tt.want {
				t.Errorf("Day8Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay8Part2(t *testing.T) {
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
			if got := Day8Part2(tt.args.input); got != tt.want {
				t.Errorf("Day8Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day8Part2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := day8Part2(tt.args.input)
			if got != tt.want {
				t.Errorf("day8Part2() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("day8Part2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
