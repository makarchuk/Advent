package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
)

type Labyrinth struct {
	scheme [][]rune
	x int
	y int
	xSpeed int
	ySpeed int
	intersection bool
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func (lab Labyrinth) getRune(x, y int) (rune, bool) {
	if y > 0 && y < len(lab.scheme) {
		if x > 0 && x < len(lab.scheme[y]) {
			return lab.scheme[y][x], true
		}
		return ' ', false
	}
	return ' ', false
}

func (lab Labyrinth) get(x, y int) rune {
	r, _ := lab.getRune(x, y)
	return r
}

func (lab *Labyrinth) Move() (rune, bool) {
	if lab.intersection {
		if lab.xSpeed == 0 {
			r, ok := lab.getRune(lab.x+1, lab.y)
			if ok && r != ' ' {
				lab.xSpeed = 1
			}
			r, ok = lab.getRune(lab.x-1, lab.y)
			if ok && r != ' ' {
				lab.xSpeed = -1
			}
			lab.ySpeed = 0
		} else if lab.ySpeed == 0 {
			r, ok := lab.getRune(lab.x, lab.y+1)
			if ok && r != ' ' {
				lab.ySpeed = 1
			}
			r, ok = lab.getRune(lab.x, lab.y-1)
			if ok && r != ' ' {
				lab.ySpeed = -1
			}
			lab.xSpeed = 0
		}
		lab.x += lab.xSpeed
		lab.y += lab.ySpeed
		lab.intersection = false
	} else {
		lab.x += lab.xSpeed
		lab.y += lab.ySpeed
		if lab.get(lab.x, lab.y) == '+' {
			lab.intersection = true
		}
		if lab.get(lab.x, lab.y) == ' ' {
			return ' ', true
		}
	}
	return lab.get(lab.x, lab.y), false
}

func getLines() []string {
	lines := make([]string, 0, 0)
	file, _ := os.Open("day19/input")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		lines = append(lines, line)
		if err != nil {
			if err != io.EOF {
				fmt.Printf(" > Failed!: %v\n", err)
			}
			break
		}
	}
	return lines
}

func realMap() Labyrinth {
	var sc [][]rune
	for _, l := range getLines() {
		var row []rune
		for _, c := range l {
			row = append(row, c)
		}
		sc = append(sc, row)
	}
	entry := 0
	for i, c := range sc[0] {
		if c == '|' {
			entry = i
		}
	}
	return Labyrinth{
		sc,
		entry,
		0,
		0,
		1,
		false,
	}
}

func testMap() Labyrinth {
	scheme := [][]rune{
		{' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', '+', '-', '-', '+', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', 'A', ' ', ' ', '|', ' ', ' ', 'C', ' ', ' ', ' '},
		{' ', 'F', '-', '-', '-', '|', '-', '-', '-', '-', 'E', '|', '-', '-', '+'},
		{' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', '|', ' ', ' ', '|', ' ', ' ', 'D'},
		{' ', ' ', ' ', ' ', ' ', '+', 'B', '-', '+', ' ', ' ', '+', '-', '-', '+'},
	}
	return Labyrinth{
		scheme,
		5,
		0,
		0,
		1,
		false,
	}
}

func part1(l Labyrinth) {
	var letters []rune
	for true {
		letter, finished := l.Move()
		if letter != '|' && letter != '-' && letter != '+' && letter != ' ' {
			letters = append(letters, letter)
		}
		if finished {
			break
		}
	}
	fmt.Printf("%q\n", letters)
}

func part2(l Labyrinth) {
	steps := 0
	for true {
		_, finished := l.Move()
		steps ++
		if finished {
			break
		}
	}
	fmt.Printf("Labyrinth passed in %v steps \n", steps)
}

func main() {
	part1(testMap())
	part1(realMap())
	part2(testMap())
	part2(realMap())
}
