package cmd

import "testing"

func TestDay18p1v0(t *testing.T) {
	d := []string{
		"1 + 2 * 3 + 4 * 5 + 6",
	}
	if got := Day18Part1(d); got != 71 {
		t.Errorf("Day18Part1() = %v", got)
	}
}

func TestDay18p1v1(t *testing.T) {
	d := []string{
		"1 + (2 * 3) + (4 * (5 + 6))",
	}
	if got := Day18Part1(d); got != 51 {
		t.Errorf("Day18Part1() = %v", got)
	}
}

func TestDay18p1v2(t *testing.T) {
	d := []string{
		"2 * 3 + (4 * 5)",
	}
	if got := Day18Part1(d); got != 26 {
		t.Errorf("Day18Part1() = %v", got)
	}
}

func TestDay18p2v0(t *testing.T) {
	d := []string{"1 + 2 * 3 + 4 * 5 + 6"}
	if got := Day18Part2(d); got != 231 {
		t.Errorf("Day18Part2() = %v", got)
	}
}
func TestDay18p2v1(t *testing.T) {
	d := []string{"2 * 3 + (4 * 5)"}
	if got := Day18Part2(d); got != 46 {
		t.Errorf("Day18Part2() = %v", got)
	}
}

func TestDay18p2v2(t *testing.T) {
	d := []string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}
	if got := Day18Part2(d); got != 1445 {
		t.Errorf("Day18Part2() = %v", got)
	}
}

func TestDay18p2v3(t *testing.T) {
	d := []string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}
	if got := Day18Part2(d); got != 669060 {
		t.Errorf("Day18Part2() = %v", got)
	}
}

func TestDay18p2v4(t *testing.T) {
	d := []string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}
	if got := Day18Part2(d); got != 23340 {
		t.Errorf("Day18Part2() = %v", got)
	}
}
