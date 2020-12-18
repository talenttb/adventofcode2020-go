package cmd

import (
	"testing"
)

func TestDay2p1(t *testing.T) {
	d := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	x := Day2(d)
	t.Logf("%v", x)
}

func TestDay2p2(t *testing.T) {
	d := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	x := Day2Part2(d)
	t.Logf("%v", x)
}

func TestDay2(t *testing.T) {
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
			if got := Day2(tt.args.arr); got != tt.want {
				t.Errorf("Day2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay2Part2(t *testing.T) {
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
			if got := Day2Part2(tt.args.arr); got != tt.want {
				t.Errorf("Day2Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
