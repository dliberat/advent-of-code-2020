package main

import "testing"

func TestSetCellNoNeighbors(t *testing.T) {
	cm := makeCellMap(false)
	cm.setCell(0, 0, 0, 0)

	if len(cm.cells) != 3*3*3 {
		t.Error("One cell should be surrounded in a 3x3x3 grid.")
	}

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				loc := coord{x, y, z, 0}

				if x == 0 && y == 0 && z == 0 {
					if cm.cells[loc] != 1 {
						t.Errorf("Cell %s should be on with no neighbors.", loc.toString())
					}
				} else {
					if (cm.cells[loc] >> 1) != 1 {
						t.Errorf("Cell %s should have a single neighbor at (0,0,0)", loc.toString())
					}
				}
			}
		}
	}
}

func TestSetCellX(t *testing.T) {
	/*
			z = 0
			        x
			    -1 0  1  2
			-1	-  -  -  -
		 y   0	-  X  X  -
			 1  -  -  -  -
	*/

	cm := makeCellMap(false)
	cm.setCell(0, 0, 0, 0)
	cm.setCell(1, 0, 0, 0)

	if len(cm.cells) != 4*3*3 {
		t.Error("Two cells should be surrounded in a 4x3x3 grid.")
	}

	for x := -1; x <= 2; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				loc := coord{x: x, y: y, z: z}

				if (x == 0 && y == 0 && z == 0) || (x == 1 && y == 0 && z == 0) {
					if cm.cells[loc] != 3 {
						t.Errorf("Cell %s should be on 1 neighbor.", loc.toString())
					}
				} else if x == -1 || x == 2 {
					if (cm.cells[loc] >> 1) != 1 {
						t.Errorf("Cell %s should have a single neighbor", loc.toString())
					}
				} else {
					if (cm.cells[loc] >> 1) != 2 {
						t.Errorf("Cell %s should have two neighbors", loc.toString())
					}
				}
			}
		}
	}
}

func TestSetCellSpacing(t *testing.T) {
	cm := makeCellMap(false)
	cm.setCell(0, 0, -5, 0)
	cm.setCell(0, 0, 5, 0)
	// empty cells between the populated regions should not be in map
	if len(cm.cells) != 2*(3*3*3) {
		t.Errorf("Expected 2*3*3*3 cells but got %d", len(cm.cells))
	}
}

func TestClearCellEmpty(t *testing.T) {
	cm := makeCellMap(false)
	cm.setCell(0, 0, 0, 0)
	cm.clearCell(0, 0, 0, 0)
	if len(cm.cells) != 0 {
		t.Errorf("cellmap should be empty but found %d cells", len(cm.cells))
	}
}

func TestClearCellY(t *testing.T) {
	cm := makeCellMap(false)
	cm.setCell(0, 0, 0, 0)
	cm.setCell(0, -1, 0, 0)

	if len(cm.cells) != 3*4*3 {
		t.Error("One cell should be surrounded in a 3x4x3 grid.")
	}

	cm.clearCell(0, -1, 0, 0)

	if len(cm.cells) != 3*3*3 {
		t.Error("One cell should be surrounded in a 3x3x3 grid.")
	}

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				loc := coord{x: x, y: y, z: z}

				if x == 0 && y == 0 && z == 0 {
					if cm.cells[loc] != 1 {
						t.Errorf("Cell %s should be on with no neighbors.", loc.toString())
					}
				} else {
					if (cm.cells[loc] >> 1) != 1 {
						t.Errorf("Cell %s should have a single neighbor at (0,0,0)", loc.toString())
					}
				}
			}
		}
	}
}

func TestClearCellSpacing(t *testing.T) {
	/*
		z = 0
			- - - -
			- X X -
			- - - -
		z = 1
			- - - -
			- X X -
			- - - -
		z = 2
			- - - -
			- - - -
			- - - -
		z = 3
			- - - -
			- X X -
			- - - -

		Every node in the range -1 <= x <= 2; -1 <= y <= 1; -1 <= z <= 4 has neighbors.
		However, if we delete the cells at (0,0,3), (0,0,1), (1,0,1) and (1,0,0),
		then the two remaining cells will be isolated.
	*/
	cm := makeCellMap(false)
	cm.setCell(0, 0, 0, 0)
	cm.setCell(1, 0, 0, 0)
	cm.setCell(0, 0, 1, 0)
	cm.setCell(1, 0, 1, 0)
	cm.setCell(1, 0, 3, 0)
	cm.setCell(0, 0, 3, 0)

	if len(cm.cells) != 4*3*6 {
		t.Errorf("Expected 4*3*6 cells but got %d", len(cm.cells))
	}

	cm.clearCell(0, 0, 1, 0)
	cm.clearCell(0, 0, 3, 0)
	cm.clearCell(1, 0, 1, 0)
	cm.clearCell(1, 0, 0, 0)
	if len(cm.cells) != 2*(3*3*3) {
		t.Errorf("Expected two isolated 3x3x3 regions of cells but got %d", len(cm.cells))
	}
}

func TestInput2CellMap(t *testing.T) {
	txt := "...\n.#.\n..."
	cm := input2cellmap(txt, false)

	if len(cm.cells) != 3*3*3 {
		t.Error("One cell should be surrounded in a 3x3x3 grid.")
	}

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				loc := coord{x: x, y: y, z: z}

				if x == 0 && y == 0 && z == 0 {
					if cm.cells[loc] != 1 {
						t.Errorf("Cell %s should be on with no neighbors.", loc.toString())
					}
				} else {
					if (cm.cells[loc] >> 1) != 1 {
						t.Errorf("Cell %s should have a single neighbor at (0,0,0)", loc.toString())
					}
				}
			}
		}
	}
}

func TestCountAfterSixIters(t *testing.T) {
	txt := ".#.\n..#\n###"
	cm := input2cellmap(txt, false)

	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()

	total := cm.countActive()
	if total != 112 {
		t.Errorf("Expected 112 active cubes but counted %d", total)
	}
}

func TestCountAfterSixIters4d(t *testing.T) {
	txt := ".#.\n..#\n###"
	cm := input2cellmap(txt, true)

	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()

	total := cm.countActive()
	if total != 848 {
		t.Errorf("Expected 848 active cubes but counted %d", total)
	}
}
