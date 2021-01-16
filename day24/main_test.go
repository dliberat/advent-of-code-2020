package main

import "testing"

func TestMove01(t *testing.T) {
	txt := "esew"
	tile := move(txt, coord{0, 0})
	if tile.x != 0 || tile.y != 1 {
		t.Errorf("Expected (0, 1, 0), but got (%v)", tile)
	}
}

func TestMove02(t *testing.T) {
	txt := "nwwswee"
	tile := move(txt, coord{0, 0})
	if tile.x != 0 || tile.y != 0 {
		t.Errorf("Expected (0, 0), but got (%v)", tile)
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

func TestInitializeTileset(t *testing.T) {
	txt := "ew" // only center tile is black
	tileset := initializeTileset(txt)
	locations := []coord{
		coord{0, 0},
		coord{1, 0},
		coord{-1, 0},
		coord{0, 1},
		coord{-1, 1},
		coord{1, -1},
		coord{0, -1},
	}
	// only center tile is black
	colors := []int{1, 0, 0, 0, 0, 0, 0}
	// six surrounding tiles have center neighbor
	neighbors := []int{0, 1, 1, 1, 1, 1, 1}
	for i, loc := range locations {
		tile, ok := tileset[loc]
		if !ok {
			t.Errorf("Coord %v should be in the tileset", loc)
		}
		if tile%2 != colors[i] {
			t.Errorf("Tile at coord %v should have color %d", loc, colors[i])
		}
		if tile>>1 != neighbors[i] {
			t.Errorf("Tile at coord %v should have %d black neighbors, not %d", loc, neighbors[i], tile>>1)
		}
	}
}

func TestInitializeTileset02(t *testing.T) {
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

	tileset := initializeTileset(txt)
	count := countBlackTiles(&tileset)
	if count != 10 {
		t.Errorf("10 != %d", count)
	}
}

func TestPart2(t *testing.T) {
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

	res := part2(txt)
	if res != 2208 {
		t.Errorf("Expected 2208 black tiles but got %d", res)
	}
}
