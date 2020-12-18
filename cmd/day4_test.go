package cmd

import (
	"testing"
)

func TestDay4p1(t *testing.T) {
	d := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}
	x := Day4Part1(d)
	t.Logf("Result: %v\n", x)
}

func TestDay4p2(t *testing.T) {
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

func TestDay4p2v2(t *testing.T) {
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

func TestDay4Part1(t *testing.T) {
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
			if got := Day4Part1(tt.args.arr); got != tt.want {
				t.Errorf("Day4Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPp(t *testing.T) {
	type args struct {
		passport []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPp(tt.args.passport); got != tt.want {
				t.Errorf("checkPp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay4Part2(t *testing.T) {
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
			if got := Day4Part2(tt.args.arr); got != tt.want {
				t.Errorf("Day4Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPpv2(t *testing.T) {
	type args struct {
		passport []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPpv2(tt.args.passport); got != tt.want {
				t.Errorf("checkPpv2() = %v, want %v", got, tt.want)
			}
		})
	}
}
