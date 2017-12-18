package main

import "fmt"

type Spinlock struct {
	position int
	buffer []int
	counter int
	stepSize int
}

func (s *Spinlock) Spin() {
	insertionIndex := (s.position + s.stepSize)%len(s.buffer) + 1
	s.counter += 1
	//fmt.Printf("Buffer: %v\n", s.buffer)
	//fmt.Printf("Index: %v\n", insertionIndex-1)
	prefix := s.buffer[0:insertionIndex-1]
	suffix := s.buffer[insertionIndex-1:]
	s.buffer = append(prefix, append([]int{s.counter}, suffix...)...)
	s.position = insertionIndex
}

func testLock() Spinlock {
	return Spinlock{
		0,
		[]int{0},
		0,
		3,
	}
}

func realLock() Spinlock {
	return Spinlock{
		0,
		make([]int, 1, 50000000),
		0,
		301,
	}
}

func part1(s Spinlock) {
	for i:=0; i<2017; i++ {
		s.Spin()
	}
	fmt.Printf("Number after 2017 is %v\n", s.buffer[s.position%len(s.buffer)])
}

type SuperFastSpinlock struct {
	len int
	second int
	stepSize int
	position int
	counter int
}

func (s *SuperFastSpinlock)Spin() {
	s.counter += 1
	newPosition := (s.position + s.stepSize) % s.len + 1
	if newPosition == 1 {
		s.second = s.counter
	}
	s.len += 1
	s.position = newPosition
}

func part2() {
	s := SuperFastSpinlock{
		1,
		0,
		301,
		0,
		0,
	}
	for i:=0; i<50000000; i++ {
		s.Spin()
	}
	fmt.Printf("Number after 0 is %v\n", s.second)
}

func main() {
	part1(testLock())
	part1(realLock())
	part2()
}
