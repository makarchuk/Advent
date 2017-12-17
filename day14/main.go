package main


import (
	"bytes"
	"fmt"
	"strings"
	"strconv"
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
	State           []int
	CurrentPosition int
	SkipSize        int
}

func (hash *Hash) Rotate(size int) {
	finalIndex := hash.CurrentPosition + size - 1
	for i := 0; i <= size/2-1; i++ {
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
	for i := 0; i < 64; i++ {
		hash.Round(lens)
	}
}

func (hash *Hash) result() string {
	res := ""
	for i := 0; i < 16; i++ {
		stepRes := 0
		for j := 0; j < 16; j++ {
			index := 16*i + j
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

func (hash *Hash) Binary(data string) string {
	lens := make([]int, 0, 0)
	for _, char := range data {
		lens = append(lens, int(char%256))
	}
	lens = append(lens, 17, 31, 73, 47, 23)
	hash.sparseHash(lens)
	binary := ""
	for i := 0; i < 16; i++ {
		stepRes := 0
		for j := 0; j < 16; j++ {
			index := 16*i + j
			stepRes = stepRes ^ hash.State[index]
		}

		active := strconv.FormatInt(int64(stepRes), 2)
		binary += strings.Repeat("0", 8 - len(active)) + active
	}
	return binary
}

func (hash *Hash) ActiveBits(input string) int {
	return strings.Count(hash.Binary(input), "1")
}

func totalActiveBits(input string) {
	total := 0
	for i:=0; i<128; i++ {
		row := input + "-" + strconv.Itoa(i)
		hash := NewHash()
		rowActiveBits := hash.ActiveBits(row)
		//fmt.Printf("string %v, activeBits: %v\n", row, rowActiveBits)
		total += rowActiveBits
	}
	fmt.Printf("Total Active Bits %v\n", total)
}

func buildMemory(input string) [][]bool {
	memory := make([][]bool, 128, 128)
	for rowNum:=0; rowNum<128; rowNum++ {
		memoryRow := make([]bool, 128, 128)
		row := input + "-" + strconv.Itoa(rowNum)
		hash := NewHash()
		rowHash := hash.Binary(row)
		for i, bit := range rowHash {
			memoryRow[i] = bit == '1'
		}
		memory[rowNum] = memoryRow
	}
	return memory
}

func RegionsNumber(input string) {
	memory := buildMemory(input)
	regionMap := make([][]int, 128, 128)
	for z := range regionMap {
		regionMap[z] = make([]int, 128, 128)
	}
	region := 1
	for i, row := range memory {
		for j := range row {
			region = MarkRegion(memory, regionMap, i, j, region)
		}
	}

	// Region -1 since we increment `region` after finding previous one
	fmt.Printf("TotalRegions %v\n", region-1)
}

func MarkRegion(memory [][]bool, regionMap [][]int, i, j, region int) int {
	if i>=0 && j>=0 && i < len(memory) && j < len(memory[i]) {
		if memory[i][j] && regionMap[i][j] == 0 {
			regionMap[i][j] = region
			MarkRegion(memory, regionMap, i, j+1, region)
			MarkRegion(memory, regionMap, i, j-1, region)
			MarkRegion(memory, regionMap, i-1, j, region)
			MarkRegion(memory, regionMap, i+1, j, region)
			return region + 1
		}
	}
	return region
}

func part1Test() {
	totalActiveBits("flqrgnkx")
}

func part1() {
	totalActiveBits("jxqlasbh")
}

func part2Test() {
	RegionsNumber("flqrgnkx")
}

func part2() {
	RegionsNumber("jxqlasbh")
}

func main() {
	//part1Test()
	//part1()
	part2Test()
	part2()
}
