package cmd

import (
	"testing"
)

func TestDay10p1(t *testing.T) {
	d := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	x := Day10Part1(d)
	t.Logf("Result: %v\n", x)
	if x != 220 {
		t.Fail()
	}
}

func TestDay10p2v1(t *testing.T) {
	d := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	x := Day10Part2(d)
	t.Logf("Result: %v\n", x)
	if x != 19208 {
		t.Fail()
	}
}

func TestDay10p2v2(t *testing.T) {
	d := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}
	x := Day10Part2(d)
	t.Logf("Result: %v\n", x)
	if x != 8 {
		t.Fail()
	}
}

func TestDay10Part1(t *testing.T) {
	type args struct {
		input []int
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
			if got := Day10Part1(tt.args.input); got != tt.want {
				t.Errorf("Day10Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay10Part2(t *testing.T) {
	type args struct {
		input []int
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
			if got := Day10Part2(tt.args.input); got != tt.want {
				t.Errorf("Day10Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
