package cmd

import "testing"

func TestDay23p1(t *testing.T) {
	d := []string{
		"389125467",
	}
	if got := Day23Part1(d); got != 67384529 {
		t.Errorf("Day23Part1() = %v", got)
	}
}

func TestDay23p2(t *testing.T) {
	d := []string{}
	if got := Day23Part2(d); got != 0 {
		t.Errorf("Day23Part2() = %v", got)
	}
}
