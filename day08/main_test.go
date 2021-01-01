package main

import "testing"

func TestCpuGetNextPc(t *testing.T) {
	prog := []instruction{
		instruction{address: 0, op: "acc", arg: 2},
		instruction{address: 1, op: "jmp", arg: 2},
		instruction{address: 2, op: "nop", arg: 0},
		instruction{address: 3, op: "jmp", arg: -3},
	}
	cpu := computer{pc: 0, prog: prog}
	next := cpu.getNextPc()
	if next != 1 {
		t.Errorf("Next instruction should be 1, but got %d", next)
		return
	}
	cpu.advancePc()
	if cpu.pc != 1 {
		t.Errorf("Pc should have been advanced to 1, but got %d", cpu.pc)
		return
	}

	next = cpu.getNextPc()
	if next != 3 {
		t.Errorf("Next instruction should be 3 but got %d", next)
		return
	}

	cpu.advancePc()
	if cpu.pc != 3 {
		t.Errorf("Pc should have been returned to 3, but got %d", cpu.pc)
		return
	}

	next = cpu.getNextPc()
	if next != 0 {
		t.Errorf("Next instruction should be 0 but got %d", next)
	}
}

func TestCpuAccumulate(t *testing.T) {
	prog := []instruction{
		instruction{address: 0, op: "acc", arg: 2},
		instruction{address: 1, op: "jmp", arg: 2},
		instruction{address: 2, op: "nop", arg: 0},
		instruction{address: 3, op: "acc", arg: -1},
		instruction{address: 4, op: "acc", arg: 3},
	}
	cpu := computer{pc: 0, prog: prog}
	for cpu.pc < len(cpu.prog) {
		cpu.executeCurrentInstr()
		cpu.advancePc()
	}
	if cpu.accumulator != 4 {
		t.Errorf("Expected accumulator=4 but got %d", cpu.accumulator)
	}
}

func TestPart2Integration01(t *testing.T) {
	prog := []instruction{
		instruction{address: 0, op: "nop", arg: 0},
		instruction{address: 1, op: "acc", arg: 1},
		instruction{address: 2, op: "jmp", arg: 4},
		instruction{address: 3, op: "acc", arg: 3},
		instruction{address: 4, op: "jmp", arg: -3},
		instruction{address: 5, op: "acc", arg: -99},
		instruction{address: 6, op: "acc", arg: 1},
		instruction{address: 7, op: "jmp", arg: -4},
		instruction{address: 8, op: "acc", arg: 6},
	}
	result, err := part2(prog)
	if err != nil {
		t.Errorf("Failed to find solution. %v", err)
		return
	}
	if result != 8 {
		t.Errorf("%d != 8", result)
	}

}
