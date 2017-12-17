package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Generator struct {
	factor int
	divisor int
	previous int
	criteria int
}

func (g *Generator) yield() int {
	value := g.previous * g.factor % g.divisor
	g.previous = value
	return value
}

func (g *Generator) yieldMatching() int {
	for true {
		value := g.yield()
		if value % g.criteria == 0 {
			return  value
		}
	}
	return 0
}

func GeneratorA(seed int) Generator {
	return Generator{
		16807,
		2147483647,
		seed,
		4,
	}
}

func GeneratorB(seed int) Generator {
	return Generator{
		48271,
		2147483647,
		seed,
		8,
	}
}

type Judge struct {
	a Generator
	b Generator
	score int
}

func (j Judge) ScoreAfter(steps int) int {
	for i:=0; i<steps; i++  {
		a, b := j.a.yield(), j.b.yield()
		if lower16bits(a) == lower16bits(b) {
			j.score += 1
		}
	}
	return j.score
}

func (j Judge) MatchingScoreAfter(steps int) int {
	for i:=0; i<steps; i++  {
		a, b := j.a.yieldMatching(), j.b.yieldMatching()
		if lower16bits(a) == lower16bits(b) {
			j.score += 1
		}
	}
	return j.score
}


func lower16bits(x int) string {
	binary := strconv.FormatInt(int64(x), 2)
	binary = strings.Repeat("0", 16) + binary
	return binary[len(binary)-16:]
}

func testJudge() Judge {
	return Judge{
		GeneratorA(65),
		GeneratorB(8921),
		0,
	}
}

func realJudge() Judge {
	return Judge{
		GeneratorA(116),
		GeneratorB(299),
		0,
	}
}

func part1Test() {
	j := testJudge()
	steps := 5
	fmt.Printf("Score after %v is %v\n", steps, j.ScoreAfter(steps))
}

func part1() {
	j := realJudge()
	steps := 40*1000*1000
	fmt.Printf("Score after %v is %v\n", steps, j.ScoreAfter(steps))
}

func part2Test() {
	j := testJudge()
	steps := 5 *1000 * 1000
	fmt.Printf("Score after %v is %v\n", steps, j.MatchingScoreAfter(steps))
}

func part2() {
	j := realJudge()
	steps := 5 *1000 * 1000
	fmt.Printf("Score after %v is %v\n", steps, j.MatchingScoreAfter(steps))
}

func main() {
	part1Test()
	part1()
	part2Test()
	part2()
}
