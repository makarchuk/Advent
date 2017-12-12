package main

import (
	"strings"
	"strconv"
	"fmt"
	"os"
	"bufio"
	"io"
)

type ProcGraph [][]bool

func NewGraph(size int) ProcGraph {
	pg := make(ProcGraph, size, size)
	for i := range pg {
		pg[i] = make([]bool, size, size)
	}
	return pg
}

func (pg ProcGraph) addPath(from, to int) {
	pg[from][to] = true
}

func (pg ProcGraph) parseRow(row string) {
	chunks := strings.Split(row, " <-> ")
	from, _ := strconv.Atoi(chunks[0])
	for _, dst := range strings.Split(chunks[1], ", ") {
		to, _ := strconv.Atoi(dst)
		pg.addPath(from, to)
	}
}

func (pg ProcGraph) availableProcesses(pid int) []int {
	available := make([]bool, len(pg), len(pg))
	available = pg.availableProcessesExcept(pid, available)
	res := make([]int, 0, 0)
	for i, avail := range available {
		if avail {
			res = append(res, i)
		}
	}
	return res
}

func (pg ProcGraph) groupsCount() int {
	groups := 0
	excluded := make([]int, len(pg), len(pg))
	for i := range pg {
		if !includes(excluded, i) {
			groups += 1
			excluded = append(excluded, pg.availableProcesses(i)...)
		}
	}
	return groups
}

func includes(list []int, num int) bool {
	for _, el := range list {
		if num == el {
			return true
		}
	}
	return false
}

func (pg ProcGraph) availableProcessesExcept(pid int, known []bool) []bool {
	known[pid] = true
	for i, avail := range pg[pid] {
		if avail {
			wasTrue := known[i]
			known[i] = true
			if !wasTrue {
				childAvailables := pg.availableProcessesExcept(i, known)
				for j, childAvail := range childAvailables {
					if childAvail {
						known[j] = childAvail
					}
				}
			}
		}
	}
	return known
}

func (pg ProcGraph) availableCount(pid int) int {
	return len(pg.availableProcesses(pid))
}

func parseInput(rows []string) ProcGraph {
	pg := NewGraph(len(rows))
	for _, row := range rows {
		pg.parseRow(row)
	}
	return pg
}

func part1Test() {
	input := []string{
		"0 <-> 2",
		"1 <-> 1",
		"2 <-> 0, 3, 4",
		"3 <-> 2, 4",
		"4 <-> 2, 3, 6",
		"5 <-> 6",
		"6 <-> 4, 5",
	}
	pg := parseInput(input)
	fmt.Printf("From root %v processes available", pg.availableCount(0))
}

func part2Test() {
	input := []string{
		"0 <-> 2",
		"1 <-> 1",
		"2 <-> 0, 3, 4",
		"3 <-> 2, 4",
		"4 <-> 2, 3, 6",
		"5 <-> 6",
		"6 <-> 4, 5",
	}
	pg := parseInput(input)
	fmt.Printf("Groups count is %v\n", pg.groupsCount())
}

func realInput() []string {
	lines := make([]string, 0, 0)
	file, _ := os.Open("day12/input")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		lines = append(lines, line)
		if err != nil {
			if err != io.EOF {
				println(" > Failed!: %v\n", err)
			}
			break
		}
	}
	return lines
}

func part1() {
	pg := parseInput(realInput())
	fmt.Printf("From root %v processes available", pg.availableCount(0))
}

func part2() {
	pg := parseInput(realInput())
	fmt.Printf("Groups count is %v\n", pg.groupsCount())
}

func main() {
	//part1()
	//part2Test()
	part2()
}
