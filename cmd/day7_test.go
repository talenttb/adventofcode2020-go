package cmd

import (
	"testing"
)

func TestDay7p1(t *testing.T) {
	d := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}
	x := Day7Part1(d)
	t.Logf("Result: %v\n", x)
	if x != 4 {
		t.Fail()
	}
}

func TestDay7p2v1(t *testing.T) {
	d := []string{
		"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
	}
	x := Day7Part2(d)
	t.Logf("Result: %v\n", x)
	if x != 126 {
		t.Fail()
	}
}

func TestDay7p2v2(t *testing.T) {
	d := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}
	x := Day7Part2(d)
	t.Logf("Result: %v\n", x)
	if x != 32 {
		t.Fail()
	}
}

func TestDay7Part1(t *testing.T) {
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
			if got := Day7Part1(tt.args.input); got != tt.want {
				t.Errorf("Day7Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay7Part2(t *testing.T) {
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
			if got := Day7Part2(tt.args.input); got != tt.want {
				t.Errorf("Day7Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myGet(t *testing.T) {
	type args struct {
		r map[string][]node
		k string
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
			if got := myGet(tt.args.r, tt.args.k); got != tt.want {
				t.Errorf("myGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
