package cmd

import (
	"testing"
)

func TestDay9p1(t *testing.T) {
	d := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	x := Day9Part1(d, 5)
	t.Logf("Result: %v\n", x)
	if x != 127 {
		t.Fail()
	}
}

func TestDay9p2v2(t *testing.T) {
	d := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	x := Day9Part2(d, 5)
	t.Logf("Result: %v\n", x)
	if x != 62 {
		t.Fail()
	}
}

func TestDay9Part1(t *testing.T) {
	type args struct {
		input      []int
		preambleNO int
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
			if got := Day9Part1(tt.args.input, tt.args.preambleNO); got != tt.want {
				t.Errorf("Day9Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay9Part2(t *testing.T) {
	type args struct {
		input      []int
		preambleNO int
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
			if got := Day9Part2(tt.args.input, tt.args.preambleNO); got != tt.want {
				t.Errorf("Day9Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		xi []int
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
			if got := sum(tt.args.xi...); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
