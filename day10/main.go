package main

import (
	"fmt"
	"bytes"
)


func upTo255() []int {
	res := make([]int, 256, 256)
	for i := range res {
		res[i] = i
	}
	return res
}

func NewHash() Hash {
	return Hash{
		upTo255(),
		0,
		0,
	}
}

type Hash struct {
	State []int
	CurrentPosition int
	SkipSize int
}

func (hash *Hash) Rotate (size int) {
	finalIndex := hash.CurrentPosition + size - 1
	for i := 0; i <= size / 2 - 1; i++ {
		first := (hash.CurrentPosition + i) % len(hash.State)
		last := (finalIndex - i) % len(hash.State)
		//fmt.Printf("First %v, Last %v\n", first, last)
		hash.State[first], hash.State[last] = hash.State[last], hash.State[first]
	}
	hash.CurrentPosition = (hash.CurrentPosition + hash.SkipSize + size) % len(hash.State)
	hash.SkipSize += 1
}

func (hash *Hash) Round(lens []int) int {
	for _, l := range lens {
		hash.Rotate(l)
	}
	return hash.State[0] * hash.State[1]
}

func (hash *Hash) sparseHash(lens []int) {
	for i:=0; i<64; i++  {
		hash.Round(lens)
	}
}

func (hash *Hash) result() string {
	res := ""
	for i:=0; i<16; i++ {
		stepRes := 0
		for j:=0; j<16; j++ {
			index := 16 * i + j
			stepRes = stepRes ^ hash.State[index]
		}
		resBuffer := bytes.NewBuffer(make([]byte, 0, 2))
		if stepRes > 15 {
			fmt.Fprintf(resBuffer, "%x", stepRes)
		} else {
			fmt.Fprintf(resBuffer, "0%x", stepRes)
		}
		res += resBuffer.String()
	}
	return res
}

func (hash *Hash) HexDigest(data string) string {
	lens := make([]int, 0, 0)
	for _, char := range data {
		lens = append(lens, int(char%256))
	}
	lens = append(lens, 17, 31, 73, 47, 23)
	hash.sparseHash(lens)
	return hash.result()
}

func part1() {
	lens := []int{76,1,88,148,166,217,130,0,128,254,16,2,130,71,255,229}
	state := Hash{
		upTo255(),
		0,
		0,
	}
	fmt.Printf("Hash is %v\n", state.Round(lens))
}

func part1Test() {
	originalState := []int{0, 1, 2, 3, 4}
	state := Hash{
		originalState,
		0,
		0,
	}
	fmt.Printf("Hash is %v\n", state.Round([]int{3, 4, 1, 5}))
}

func part2Test() {
	for _, inp := range []string{"", "AoC 2017", "1,2,3", "1,2,4"} {
		hash := NewHash()
		fmt.Printf("Hash of %v is %v\n", inp, hash.HexDigest(inp))
	}
}

func part2() {
	hash := NewHash()
	fmt.Printf("Hash is %v\n", hash.HexDigest("76,1,88,148,166,217,130,0,128,254,16,2,130,71,255,229"))
}

func main() {
	part2Test()
	part2()
}
