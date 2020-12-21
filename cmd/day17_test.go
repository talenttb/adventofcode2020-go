package cmd

import "testing"

func TestDay17p1(t *testing.T) {
	d := []string{
		".#.",
		"..#",
		"###",
	}
	if got := Day17Part1(d); got != 0 {
		t.Errorf("Day17Part1() = %v", got)
	}
}

func TestDay17p2(t *testing.T) {
	d := []string{}
	if got := Day17Part2(d); got != 0 {
		t.Errorf("Day17Part2() = %v", got)
	}
}
