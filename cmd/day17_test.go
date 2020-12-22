package cmd

import "testing"

func TestDay17p1v0(t *testing.T) {
	d := []string{
		".#.",
		"..#",
		"###",
	}
	if got := Day17Part1(d, 1); got != 11 {
		t.Errorf("Day17Part1() = %v", got)
	}
}

func TestDay17p1v1(t *testing.T) {
	d := []string{
		".#.",
		"..#",
		"###",
	}
	if got := Day17Part1(d, 6); got != 112 {
		t.Errorf("Day17Part1() = %v", got)
	}
}

func TestDay17p2(t *testing.T) {
	d := []string{}
	if got := Day17Part2(d, 6); got != 0 {
		t.Errorf("Day17Part2() = %v", got)
	}
}
