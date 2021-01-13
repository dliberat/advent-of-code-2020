package main

import "testing"

func TestMove01(t *testing.T) {
	txt := "esew"
	tile := move(txt, coord{0, 0, 0})
	if tile.x != 0 || tile.y != 1 || tile.z != 0 {
		t.Errorf("Expected (0, 1, 0), but got (%v)", tile)
	}
}

func TestMove02(t *testing.T) {
	txt := "nwwswee"
	tile := move(txt, coord{0, 0, 0})
	if tile.x != 0 || tile.y != 0 || tile.z != 0 {
		t.Errorf("Expected (0, 0, 0), but got (%v)", tile)
	}
}

func TestPart1(t *testing.T) {
	txt := `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

	res := part1(txt)
	if res != 10 {
		t.Errorf("Expected 10 black tiles but got %d", res)
	}
}
