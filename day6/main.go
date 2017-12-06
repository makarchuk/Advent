package main

import (
	"fmt"
)
type Memory []int

func (mem Memory) step() Memory {
	newMem := make(Memory, len(mem), len(mem))
	copy(newMem, mem)
	overflowedCell := mem.overflowed()
	for i:= 0; i < mem[overflowedCell]; i++  {
		index := (1 + i + overflowedCell) % len(mem)
		newMem[overflowedCell] -= 1
		newMem[index] += 1
	}
	return newMem
}

func (mem Memory) equals(other Memory) bool {
	for i := range mem {
		if mem[i] != other[i] {
			return false
		}
	}
	return true
}

func (mem Memory) overflowed()  int {
	currentI := 0
	for i, e := range mem {
		if e > mem[currentI] {
			currentI = i
		}
	}
	return currentI
}

func (mem Memory) in(states []Memory) bool {
	for _, m := range states {
		if mem.equals(m) {
			return true
		}
	}
	return false
}


func (mem Memory) cycleLen(states []Memory) (int, bool) {
	for i:=len(states)-1; i>=0; i--  {
		if states[i].equals(mem) {
			return len(states) - i, true
		}
	}
	return 0, false
}

func part1(m Memory) {
	states := make([]Memory, 0, 0)
	counter := 0
	for {
		newState := m.step()
		states = append(states, m)
		counter += 1
		_, match := newState.cycleLen(states)
		if match {
			break
		}
		m = newState
	}
	fmt.Printf("Match on a state #%d", counter)
}

func part2(m Memory) {
	states := make([]Memory, 0, 0)
	var cycleLen int
	counter := 0
	for {
		newState := m.step()
		states = append(states, m)
		counter += 1
		cycle, match := newState.cycleLen(states)
		cycleLen = cycle
		if match {
			break
		}
		m = newState
	}
	fmt.Printf("Cycle len is #%d", cycleLen)
}

func main() {
	input := Memory {0, 2, 7, 0}
	input = Memory {2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14}
	part2(input)
}
