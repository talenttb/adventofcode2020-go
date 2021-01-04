package cmd

import "testing"

func TestDay21p1(t *testing.T) {
	d := []string{}
	if got := Day21Part1(d); got != 0 {
		t.Errorf("Day21Part1() = %v", got)
	}
}

func TestDay21p2(t *testing.T) {
	d := []string{}
	if got := Day21Part2(d); got != 0 {
		t.Errorf("Day21Part2() = %v", got)
	}
}
