package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	address int
	op      string
	arg     int
}

type program []instruction

type computer struct {
	pc          int
	accumulator int
	prog        program
}

func (cpu *computer) getCurrentInstr() instruction {
	return cpu.prog[cpu.pc]
}
func (cpu *computer) getNextPc() int {
	currentInstr := cpu.getCurrentInstr()
	if currentInstr.op == "jmp" {
		return cpu.pc + currentInstr.arg
	}
	return cpu.pc + 1
}
func (cpu *computer) executeCurrentInstr() {
	currentInstr := cpu.getCurrentInstr()
	if currentInstr.op == "acc" {
		cpu.accumulator += currentInstr.arg
	}
}
func (cpu *computer) advancePc() {
	cpu.pc = cpu.getNextPc()
}
func (cpu *computer) isTerminated() bool {
	// The program is supposed to terminate by attempting to execute
	// an instruction immediately after the last instruction in the file.
	return cpu.pc == len(cpu.prog)
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func parseProgram(txt string) program {
	lines := strings.Split(txt, "\n")
	p := make(program, len(lines))
	for addr, line := range lines {
		split := strings.Split(line, " ")
		op, strArg := split[0], split[1]
		arg, err := strconv.Atoi(strArg)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse program at line %d", addr))
		}
		instr := instruction{address: addr, op: op, arg: arg}
		p[addr] = instr
	}
	return p
}

func part1(prog program) (int, error) {
	cpu := computer{prog: prog}
	history := make(map[int]bool, 0)
	for !cpu.isTerminated() {
		visited := history[cpu.pc]
		if visited {
			// same instruction has been executed previously
			return cpu.accumulator, nil
		}

		history[cpu.pc] = true
		cpu.executeCurrentInstr()
		cpu.advancePc()
	}
	return cpu.accumulator, errors.New("failed to detect a loop in the program")
}

func part2(prog program) (int, error) {
	for i, instr := range prog {
		if instr.op != "jmp" && instr.op != "nop" {
			continue
		}

		clone := make(program, len(prog))
		copy(clone, prog)
		if instr.op == "jmp" {
			clone[i].op = "nop"
		} else {
			clone[i].op = "jmp"
		}
		res, err := part1(clone)
		if err != nil {
			return res, nil
		}
	}
	return 0, errors.New("all variations ended in infinite loops")
}

func main() {
	txt := readInput("input.txt")
	prog := parseProgram(txt)
	p1, err := part1(prog)
	if err != nil {
		fmt.Println("[PART 1]", err)
	} else {
		fmt.Println("[PART 1] Accumulator value before first repeated instruction:", p1)
	}
	p2, err := part2(prog)
	if err != nil {
		fmt.Println("[PART 2]", err)
	} else {
		fmt.Println("[PART 2] Result:", p2)
	}
}
