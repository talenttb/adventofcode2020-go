package cmd

import (
	"testing"
)

func TestDay11p1(t *testing.T) {
	d := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	x := Day11Part1(d)
	t.Logf("Result: %v\n", x)
	if x != 37 {
		t.Fail()
	}
}

func TestDay11p2(t *testing.T) {
	d := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	x := Day11Part2(d)
	t.Logf("Result: %v\n", x)
	if x != 26 {
		t.Fail()
	}
}

func TestDay11Part1(t *testing.T) {
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
			if got := Day11Part1(tt.args.input); got != tt.want {
				t.Errorf("Day11Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay11Part2(t *testing.T) {
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
			if got := Day11Part2(tt.args.input); got != tt.want {
				t.Errorf("Day11Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
