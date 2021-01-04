package cmd

import "testing"

func TestDay22p1(t *testing.T) {
	d := []string{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
		"",
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	}
	if got := Day22Part1(d); got != 306 {
		t.Errorf("Day22Part1() = %v", got)
	}
}

func TestDay22p2(t *testing.T) {
	d := []string{}
	if got := Day22Part2(d); got != 0 {
		t.Errorf("Day22Part2() = %v", got)
	}
}
