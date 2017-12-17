package main

import (
	"strings"
	"strconv"
	"fmt"
)

type Layer struct {
	Range int
}

type Firewall struct {
	Layers map[int]Layer
	Depth int
	Position int
}

func ParseLayer(row string) (int, Layer) {
	chunks := strings.Split(strings.TrimSpace(row), ": ")
	layerNumber, _ := strconv.Atoi(chunks[0])
	layerRange, _ := strconv.Atoi(chunks[1])
	layer := Layer{layerRange - 1}
	return layerNumber, layer
}

func ParseFirewall(rows []string) Firewall {
	fw := Firewall{make(map[int]Layer), 0, -1}
	for _, row := range rows {
		layerNum, layer := ParseLayer(row)
		fw.Layers[layerNum] = layer
		fw.Depth = layerNum
	}
	return fw
}

func (l Layer) IsZero(step int) bool {
	return step % (2*l.Range) == 0
}

func (fw *Firewall) Step() int {
	caught := 0
	fw.Position += 1
	l, ok := fw.Layers[fw.Position]
	if ok && l.IsZero(fw.Position) {
		fmt.Println("TRIGGERED!")
		caught = fw.Position * (l.Range + 1)
	}
	return caught
}

func (fw Firewall) Walk() int {
	severity := 0
	for fw.Position <= fw.Depth {
		stepSeverity := fw.Step()
		severity += stepSeverity
	}
	fmt.Printf("Total severity is %v\n", severity)
	return severity
}

func (fw Firewall) Dodge() {
	delay := 0
	for true {
		success := true
		for depth, layer := range fw.Layers {
			if layer.IsZero(depth + delay) {
				success = false
				fmt.Printf("Caught on delay %v and layer %v\n", delay, layer)
				break
			}
		}
		if success {
			break
		}
		delay ++
	}
	fmt.Printf("Succes on delay %v\n", delay)
}

func part1() {
	fw := ParseFirewall(realInput())
	fw.Walk()
}

func part1Test() {
	fw := ParseFirewall(testInput())
	fw.Walk()
}

func testInput() []string {
	return []string {
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}
}

func realInput() []string {
	return []string {
		"0: 3",
		"1: 2",
		"2: 4",
		"4: 8",
		"6: 5",
		"8: 6",
		"10: 6",
		"12: 4",
		"14: 6",
		"16: 6",
		"18: 17",
		"20: 8",
		"22: 8",
		"24: 8",
		"26: 9",
		"28: 8",
		"30: 12",
		"32: 12",
		"34: 10",
		"36: 12",
		"38: 12",
		"40: 8",
		"42: 12",
		"44: 12",
		"46: 10",
		"48: 12",
		"50: 12",
		"52: 14",
		"54: 14",
		"56: 12",
		"58: 14",
		"60: 14",
		"62: 14",
		"64: 14",
		"66: 14",
		"68: 12",
		"70: 14",
		"72: 14",
		"74: 14",
		"76: 14",
		"80: 18",
		"82: 14",
		"90: 18",
	}
}

func part2Test() {
	fw := ParseFirewall(testInput())
	fw.Dodge()
}

func part2() {
	fw := ParseFirewall(realInput())
	fw.Dodge()
}

func main() {
	//part1Test()
	//part1()
	part2Test()
	part2()
}
