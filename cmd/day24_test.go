package cmd

import "testing"

func TestDay24p1(t *testing.T) {
	d := []string{
		"sesenwnenenewseeswwswswwnenewsewsw",
		"neeenesenwnwwswnenewnwwsewnenwseswesw",
		"seswneswswsenwwnwse",
		"nwnwneseeswswnenewneswwnewseswneseene",
		"swweswneswnenwsewnwneneseenw",
		"eesenwseswswnenwswnwnwsewwnwsene",
		"sewnenenenesenwsewnenwwwse",
		"wenwwweseeeweswwwnwwe",
		"wsweesenenewnwwnwsenewsenwwsesesenwne",
		"neeswseenwwswnwswswnw",
		"nenwswwsewswnenenewsenwsenwnesesenew",
		"enewnwewneswsewnwswenweswnenwsenwsw",
		"sweneswneswneneenwnewenewwneswswnese",
		"swwesenesewenwneswnwwneseswwne",
		"enesenwswwswneneswsenwnewswseenwsese",
		"wnwnesenesenenwwnenwsewesewsesesew",
		"nenewswnwewswnenesenwnesewesw",
		"eneswnwswnwsenenwnwnwwseeswneewsenese",
		"neswnwewnwnwseenwseesewsenwsweewe",
		"wseweeenwnesenwwwswnew",
	}
	if got := Day24Part1(d); got != 10 {
		t.Errorf("Day24Part1() = %v", got)
	}
}

func TestDay24p1v1(t *testing.T) {
	d := []string{
		"esew",
	}
	if got := Day24Part1(d); got != 1 {
		t.Errorf("Day24Part1() = %v", got)
	}
}

func TestDay24p1v2(t *testing.T) {
	d := []string{
		"nwwswee",
	}
	if got := Day24Part1(d); got != 1 {
		t.Errorf("Day24Part1() = %v", got)
	}
}

func TestDay24p1v3(t *testing.T) {
	d := []string{
		"sesenwnenenewseeswwswswwnenewsewsw",
	}
	if got := Day24Part1(d); got != 0 {
		t.Errorf("Day24Part1() = %v", got)
	}
}

func TestDay24p2(t *testing.T) {
	d := []string{}
	if got := Day24Part2(d); got != 0 {
		t.Errorf("Day24Part2() = %v", got)
	}
}
