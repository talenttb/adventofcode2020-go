package cmd

import (
	"testing"
)

func TestDay5p1(t *testing.T) {
	d := "BFFFBBFRRR"
	x := Day5Part1(d)
	t.Logf("Result: %v\n", x)
	if x != 567 {
		t.Fail()
	}
}

func TestDay5p1v1(t *testing.T) {
	d := "BBFFBBFRLL"
	x := Day5Part1(d)
	t.Logf("Result: %v\n", x)
	if x != 820 {
		t.Fail()
	}
}

func TestDay5p1v2(t *testing.T) {
	d := "FFFBBBFRRR"
	x := Day5Part1(d)
	t.Logf("Result: %v\n", x)
	if x != 119 {
		t.Fail()
	}
}

func TestDay5p2(t *testing.T) {
	d := []string{
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"",
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946",
		"",
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"",
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007",
	}
	x := Day4Part2(d)
	t.Logf("Result: %v\n", x)
}

func TestDay5p2v2(t *testing.T) {
	d := []string{
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
		"hcl:#623a2f",
		"",
		"eyr:2029 ecl:blu cid:129 byr:1989",
		"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"",
		"hcl:#888785",
		"hgt:164cm byr:2001 iyr:2015 cid:88",
		"pid:545766238 ecl:hzl",
		"eyr:2022",
		"",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	}
	x := Day4Part2(d)
	t.Logf("Result: %v\n", x)
}

func TestDay5Part1(t *testing.T) {
	type args struct {
		seatID string
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
			if got := Day5Part1(tt.args.seatID); got != tt.want {
				t.Errorf("Day5Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
